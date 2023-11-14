package firestore

import (
	"context"
	"encoding/json"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/tinode/chat/server/auth"
	t "github.com/tinode/chat/server/store/types"
)

const (
	UserCollection            = "ChatUsers"
	TagCollection             = "ChatTags"
	DeviceCollection          = "ChatDevices"
	AuthCollection            = "ChatAuth"
	TopicCollection           = "ChatTopic"
	TopicTagCollection        = "ChatTopicTags"
	SubscriptionCollection    = "ChatSubscriptions"
	MessageCollection         = "ChatMessages"
	DellogCollection          = "ChatDellog"
	CredentialCollection      = "ChatCredentials"
	FileUploadCollection      = "ChatFileUploads"
	FileMessageLinkCollection = "ChatFileMessageLinks"
)

var ctx = context.Background()

type firestoreDB struct {
	fs *firestore.Client
}

func (fs *firestoreDB) Open(config json.RawMessage) (err error) {
	fs.fs, err = firestore.NewClient(ctx, "darren-prd")
	if err != nil {
		return err
	}
	return nil
}

func (fs *firestoreDB) Close() error                { return nil }
func (fs *firestoreDB) IsOpen() bool                { return true }
func (fs *firestoreDB) GetDbVersion() (int, error)  { return 0, nil }
func (fs *firestoreDB) CheckDbVersion() error       { return nil }
func (fs *firestoreDB) GetName() string             { return "" }
func (fs *firestoreDB) SetMaxResults(val int) error { return nil }
func (fs *firestoreDB) CreateDb(reset bool) error   { return nil }
func (fs *firestoreDB) UpgradeDb() error            { return nil }
func (fs *firestoreDB) Version() int                { return 0 }
func (fs *firestoreDB) Stats() interface{}          { return nil }

func (fs *firestoreDB) UserCreate(user *t.User) error {
	_, err := fs.fs.Collection(UserCollection).Doc(user.Id).Set(ctx, user)
	return err
}

func (fs *firestoreDB) UserGet(uid t.Uid) (*t.User, error) {
	snap, err := fs.fs.Collection(UserCollection).Doc(uid.String()).Get(ctx)
	if err != nil {
		return nil, err
	}
	var user t.User
	err = snap.DataTo(&user)
	return &user, err
}

func (fs *firestoreDB) UserGetAll(ids ...t.Uid) ([]t.User, error) {
	snaps, err := fs.fs.Collection(UserCollection).Query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	var users []t.User
	for _, snap := range snaps {
		var user t.User
		err = snap.DataTo(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (fs *firestoreDB) UserDelete(uid t.Uid, hard bool) error {
	_, err := fs.fs.Collection(UserCollection).Doc(uid.String()).Delete(ctx)
	return err
}

func (fs *firestoreDB) UserUpdate(uid t.Uid, update map[string]interface{}) error {
	return nil // todo
}

func (fs *firestoreDB) UserUpdateTags(uid t.Uid, add, remove, reset []string) ([]string, error) {

}

func (fs *firestoreDB) UserGetByCred(method, value string) (t.Uid, error) {
	snaps, err := fs.fs.Collection(UserCollection).Query.Where("synthetic", "=", method+":"+value).Documents(ctx).GetAll()
	if err != nil {
		return t.ZeroUid, err
	}
	for _, snap := range snaps {
		var user t.User
		err = snap.DataTo(&user)
		if err != nil {
			return t.ZeroUid, err
		}
		return user.Uid(), nil
	}
	return t.ZeroUid, nil
}

func (fs *firestoreDB) UserUnreadCount(ids ...t.Uid) (map[t.Uid]int, error) {

}

func (fs *firestoreDB) UserGetUnvalidated(lastUpdatedBefore time.Time, limit int) ([]t.Uid, error) {

}

func (fs *firestoreDB) CredUpsert(cred *t.Credential) (bool, error) {

}

func (fs *firestoreDB) CredGetActive(uid t.Uid, method string) (*t.Credential, error) {

}

func (fs *firestoreDB) CredGetAll(uid t.Uid, method string, validatedOnly bool) ([]t.Credential, error) {

}

func (fs *firestoreDB) CredDel(uid t.Uid, method, value string) error {

}

func (fs *firestoreDB) CredConfirm(uid t.Uid, method string) error {

}

func (fs *firestoreDB) CredFail(uid t.Uid, method string) error {

}

func (fs *firestoreDB) AuthGetUniqueRecord(unique string) (t.Uid, auth.Level, []byte, time.Time, error) {

}

func (fs *firestoreDB) AuthGetRecord(user t.Uid, scheme string) (string, auth.Level, []byte, time.Time, error) {

}

func (fs *firestoreDB) AuthAddRecord(user t.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {

}

func (fs *firestoreDB) AuthDelScheme(user t.Uid, scheme string) error {

}

func (fs *firestoreDB) AuthDelAllRecords(uid t.Uid) (int, error) {

}

func (fs *firestoreDB) AuthUpdRecord(user t.Uid, scheme, unique string, authLvl auth.Level, secret []byte, expires time.Time) error {

}

func (fs *firestoreDB) TopicCreate(topic *t.Topic) error {
	_, err := fs.fs.Collection(UserCollection).Doc(topic.Id).Set(ctx, topic)
	return err
}

func (fs *firestoreDB) TopicCreateP2P(initiator, invited *t.Subscription) error {

}

func (fs *firestoreDB) TopicGet(topic string) (*t.Topic, error) {

}

func (fs *firestoreDB) TopicsForUser(uid t.Uid, keepDeleted bool, opts *t.QueryOpt) ([]t.Subscription, error) {

}

func (fs *firestoreDB) UsersForTopic(topic string, keepDeleted bool, opts *t.QueryOpt) ([]t.Subscription, error) {

}

func (fs *firestoreDB) OwnTopics(uid t.Uid) ([]string, error) {

}

func (fs *firestoreDB) ChannelsForUser(uid t.Uid) ([]string, error) {

}

func (fs *firestoreDB) TopicShare(subs []*t.Subscription) error {

}

func (fs *firestoreDB) TopicDelete(topic string, isChan, hard bool) error {

}

func (fs *firestoreDB) TopicUpdateOnMessage(topic string, msg *t.Message) error {

}

func (fs *firestoreDB) TopicUpdate(topic string, update map[string]interface{}) error {

}

func (fs *firestoreDB) TopicOwnerChange(topic string, newOwner t.Uid) error {

}

func (fs *firestoreDB) SubscriptionGet(topic string, user t.Uid, keepDeleted bool) (*t.Subscription, error) {

}

func (fs *firestoreDB) SubsForUser(user t.Uid) ([]t.Subscription, error) {

}

func (fs *firestoreDB) SubsForTopic(topic string, keepDeleted bool, opts *t.QueryOpt) ([]t.Subscription, error) {

}

func (fs *firestoreDB) SubsUpdate(topic string, user t.Uid, update map[string]interface{}) error {

}

func (fs *firestoreDB) SubsDelete(topic string, user t.Uid) error {

}

func (fs *firestoreDB) FindUsers(user t.Uid, req [][]string, opt []string, activeOnly bool) ([]t.Subscription, error) {

}

func (fs *firestoreDB) FindTopics(req [][]string, opt []string, activeOnly bool) ([]t.Subscription, error) {

}

func (fs *firestoreDB) MessageSave(msg *t.Message) error {

}

func (fs *firestoreDB) MessageGetAll(topic string, forUser t.Uid, opts *t.QueryOpt) ([]t.Message, error) {

}

func (fs *firestoreDB) MessageDeleteList(topic string, toDel *t.DelMessage) error {

}

func (fs *firestoreDB) MessageGetDeleted(topic string, forUser t.Uid, opts *t.QueryOpt) ([]t.DelMessage, error) {

}

func (fs *firestoreDB) DeviceUpsert(uid t.Uid, dev *t.DeviceDef) error {

}

func (fs *firestoreDB) DeviceGetAll(uid ...t.Uid) (map[t.Uid][]t.DeviceDef, int, error) {

}

func (fs *firestoreDB) DeviceDelete(uid t.Uid, deviceID string) error {

}

func (fs *firestoreDB) FileStartUpload(fd *t.FileDef) error {

}

func (fs *firestoreDB) FileFinishUpload(fd *t.FileDef, success bool, size int64) (*t.FileDef, error) {

}

func (fs *firestoreDB) FileGet(fid string) (*t.FileDef, error) {

}

func (fs *firestoreDB) FileDeleteUnused(olderThan time.Time, limit int) ([]string, error) {

}

func (fs *firestoreDB) FileLinkAttachments(topic string, userId, msgId t.Uid, fids []string) error {

}

func (fs *firestoreDB) PCacheGet(key string) (string, error) {

}

func (fs *firestoreDB) PCacheUpsert(key string, value string, failOnDuplicate bool) error {

}

func (fs *firestoreDB) PCacheDelete(key string) error {

}

func (fs *firestoreDB) PCacheExpire(keyPrefix string, olderThan time.Time) error {

}
