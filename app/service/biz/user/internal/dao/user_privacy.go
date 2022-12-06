// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package dao

import (
	"context"
	"fmt"

	"github.com/teamgram/marmota/pkg/hack"
	"github.com/teamgram/marmota/pkg/stores/sqlc"
	"github.com/teamgram/marmota/pkg/stores/sqlx"
	"github.com/teamgram/proto/mtproto"
	"github.com/teamgram/teamgram-server/app/service/biz/user/internal/dal/dataobject"
	"github.com/teamgram/teamgram-server/app/service/biz/user/user"

	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/mr"
)

var (
	userPrivacyKeyStatusTimestampPrefix = "user_privacy_status_timestamp"
	userPrivacyKeyChatInvitePrefix      = "user_privacy_chat_invite"
	userPrivacyKeyPhoneCallPrefix       = "user_privacy_phone_call"
	userPrivacyKeyPhoneP2PPrefix        = "user_privacy_phone_p2p"
	userPrivacyKeyForwardsPrefix        = "user_privacy_forwards"
	userPrivacyKeyProfilePhotoPrefix    = "user_privacy_profile_photo"
	userPrivacyKeyPhoneNumberPrefix     = "user_privacy_phone_number"
	userPrivacyKeyAddedByPhonePrefix    = "user_privacy_added_by_phone"
	userPrivacyKeyVoiceMessagesPrefix   = "user_privacy_voice_messages"
)

func genUserPrivacyKeyPrefix(id int64, keyType int32) string {
	switch keyType {
	case user.STATUS_TIMESTAMP:
		return genUserPrivacyKeyStatusTimestampPrefix(id)
	case user.CHAT_INVITE:
		return genUserPrivacyKeyChatInvitePrefix(id)
	case user.PHONE_CALL:
		return genUserPrivacyKeyPhoneCallPrefix(id)
	case user.PHONE_P2P:
		return genUserPrivacyKeyPhoneP2PPrefix(id)
	case user.FORWARDS:
		return genUserPrivacyKeyForwardsPrefix(id)
	case user.PROFILE_PHOTO:
		return genUserPrivacyKeyProfilePhotoPrefix(id)
	case user.PHONE_NUMBER:
		return genUserPrivacyKeyPhoneNumberPrefix(id)
	case user.ADDED_BY_PHONE:
		return genUserPrivacyKeyAddedByPhonePrefix(id)
	case user.VOICE_MESSAGES:
		return genUserPrivacyKeyVoiceMessagesPrefix(id)
	default:
		return ""
	}
}

func genUserPrivacyKeyStatusTimestampPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyStatusTimestampPrefix, id)
}

func genUserPrivacyKeyChatInvitePrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyChatInvitePrefix, id)
}

func genUserPrivacyKeyPhoneCallPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyPhoneCallPrefix, id)
}

func genUserPrivacyKeyPhoneP2PPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyPhoneP2PPrefix, id)
}

func genUserPrivacyKeyForwardsPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyForwardsPrefix, id)
}

func genUserPrivacyKeyProfilePhotoPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyProfilePhotoPrefix, id)
}

func genUserPrivacyKeyPhoneNumberPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyPhoneNumberPrefix, id)
}

func genUserPrivacyKeyAddedByPhonePrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyAddedByPhonePrefix, id)
}

func genUserPrivacyKeyVoiceMessagesPrefix(id int64) string {
	return fmt.Sprintf("%s_%d", userPrivacyKeyVoiceMessagesPrefix, id)
}

func (d *Dao) GetUserPrivacyRulesListByKeys(ctx context.Context, id int64, keys ...int32) []*user.PrivacyKeyRules {
	var (
		cacheRules []*user.PrivacyKeyRules
		// cacheKey   string
	)

	if len(keys) == 1 {
		rules, _ := d.GetUserPrivacyRules(ctx, id, keys[0])
		if rules != nil {
			cacheRules = append(cacheRules, rules)
		}
	} else if len(keys) > 1 {
		cacheRules2 := make([]*user.PrivacyKeyRules, len(keys))
		mr.ForEach(
			func(source chan<- interface{}) {
				for i := 0; i < len(keys); i++ {
					source <- i
				}
			},
			func(item interface{}) {
				idx := item.(int)
				rules, _ := d.GetUserPrivacyRules(ctx, id, keys[idx])
				if rules != nil {
					cacheRules2[idx] = rules
				}
			})
		for _, v := range cacheRules2 {
			if v != nil {
				cacheRules = append(cacheRules, v)
			}
		}
	}

	return cacheRules
}

func (d *Dao) GetUserPrivacyRules(ctx context.Context, id int64, key int32) (*user.PrivacyKeyRules, error) {
	cacheKey := genUserPrivacyKeyPrefix(id, key)
	if cacheKey == "" {
		return nil, mtproto.ErrPrivacyKeyInvalid
	}

	rules := user.MakeTLPrivacyKeyRules(&user.PrivacyKeyRules{
		Key:   key,
		Rules: nil,
	}).To_PrivacyKeyRules()
	err := d.CachedConn.QueryRow(
		ctx,
		rules,
		cacheKey,
		func(ctx context.Context, conn *sqlx.DB, v interface{}) error {
			do, err := d.UserPrivaciesDAO.SelectPrivacy(ctx, id, key)
			if err != nil {
				return err
			}
			rv := v.(*user.PrivacyKeyRules)

			if do == nil {
				// return sqlc.ErrNotFound
				rv.Rules = []*mtproto.PrivacyRule{mtproto.MakeTLPrivacyValueAllowAll(nil).To_PrivacyRule()}
			} else {
				if err = jsonx.UnmarshalFromString(do.Rules, &rv.Rules); err != nil {
					// c.Logger.Errorf("user.getPrivacy - Unmarshal PrivacyRulesData(%d)error: %v", do.Id, err)
					return err
				}
			}

			return nil
		},
	)

	if err != nil {
		if err == sqlc.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return rules, nil
}

func (d *Dao) SetUserPrivacyRules(ctx context.Context, id int64, key int32, rules []*mtproto.PrivacyRule) bool {
	cacheKey := genUserPrivacyKeyPrefix(id, key)
	if cacheKey == "" {
		return false
	}

	cacheKeys := []string{cacheKey}
	switch key {
	case user.STATUS_TIMESTAMP,
		user.PROFILE_PHOTO,
		user.PHONE_NUMBER:
		cacheKeys = append(cacheKeys, genCacheUserDataCacheKey(id))
	}

	d.CachedConn.Exec(
		ctx,
		func(ctx context.Context, conn *sqlx.DB) (int64, int64, error) {
			rulesData, _ := jsonx.Marshal(rules)
			return d.UserPrivaciesDAO.InsertOrUpdate(ctx, &dataobject.UserPrivaciesDO{
				UserId:  id,
				KeyType: key,
				Rules:   hack.String(rulesData),
			})
		},
		cacheKeys...)

	return true
}
