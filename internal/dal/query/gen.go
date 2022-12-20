// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                  db,
		AccessToken:         newAccessToken(db, opts...),
		App:                 newApp(db, opts...),
		Cast:                newCast(db, opts...),
		Character:           newCharacter(db, opts...),
		CharacterComment:    newCharacterComment(db, opts...),
		CharacterSubjects:   newCharacterSubjects(db, opts...),
		EpCollection:        newEpCollection(db, opts...),
		Episode:             newEpisode(db, opts...),
		EpisodeComment:      newEpisodeComment(db, opts...),
		Friend:              newFriend(db, opts...),
		Group:               newGroup(db, opts...),
		GroupMember:         newGroupMember(db, opts...),
		GroupTopic:          newGroupTopic(db, opts...),
		GroupTopicComment:   newGroupTopicComment(db, opts...),
		Index:               newIndex(db, opts...),
		IndexComment:        newIndexComment(db, opts...),
		IndexSubject:        newIndexSubject(db, opts...),
		Member:              newMember(db, opts...),
		Notification:        newNotification(db, opts...),
		NotificationField:   newNotificationField(db, opts...),
		OAuthClient:         newOAuthClient(db, opts...),
		Person:              newPerson(db, opts...),
		PersonComment:       newPersonComment(db, opts...),
		PersonField:         newPersonField(db, opts...),
		PersonSubjects:      newPersonSubjects(db, opts...),
		PrivateMessage:      newPrivateMessage(db, opts...),
		RevisionHistory:     newRevisionHistory(db, opts...),
		RevisionText:        newRevisionText(db, opts...),
		Subject:             newSubject(db, opts...),
		SubjectCollection:   newSubjectCollection(db, opts...),
		SubjectField:        newSubjectField(db, opts...),
		SubjectRelation:     newSubjectRelation(db, opts...),
		SubjectRevision:     newSubjectRevision(db, opts...),
		SubjectTopic:        newSubjectTopic(db, opts...),
		SubjectTopicComment: newSubjectTopicComment(db, opts...),
		TimeLine:            newTimeLine(db, opts...),
		UserGroup:           newUserGroup(db, opts...),
		WebSession:          newWebSession(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	AccessToken         accessToken
	App                 app
	Cast                cast
	Character           character
	CharacterComment    characterComment
	CharacterSubjects   characterSubjects
	EpCollection        epCollection
	Episode             episode
	EpisodeComment      episodeComment
	Friend              friend
	Group               group
	GroupMember         groupMember
	GroupTopic          groupTopic
	GroupTopicComment   groupTopicComment
	Index               index
	IndexComment        indexComment
	IndexSubject        indexSubject
	Member              member
	Notification        notification
	NotificationField   notificationField
	OAuthClient         oAuthClient
	Person              person
	PersonComment       personComment
	PersonField         personField
	PersonSubjects      personSubjects
	PrivateMessage      privateMessage
	RevisionHistory     revisionHistory
	RevisionText        revisionText
	Subject             subject
	SubjectCollection   subjectCollection
	SubjectField        subjectField
	SubjectRelation     subjectRelation
	SubjectRevision     subjectRevision
	SubjectTopic        subjectTopic
	SubjectTopicComment subjectTopicComment
	TimeLine            timeLine
	UserGroup           userGroup
	WebSession          webSession
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		AccessToken:         q.AccessToken.clone(db),
		App:                 q.App.clone(db),
		Cast:                q.Cast.clone(db),
		Character:           q.Character.clone(db),
		CharacterComment:    q.CharacterComment.clone(db),
		CharacterSubjects:   q.CharacterSubjects.clone(db),
		EpCollection:        q.EpCollection.clone(db),
		Episode:             q.Episode.clone(db),
		EpisodeComment:      q.EpisodeComment.clone(db),
		Friend:              q.Friend.clone(db),
		Group:               q.Group.clone(db),
		GroupMember:         q.GroupMember.clone(db),
		GroupTopic:          q.GroupTopic.clone(db),
		GroupTopicComment:   q.GroupTopicComment.clone(db),
		Index:               q.Index.clone(db),
		IndexComment:        q.IndexComment.clone(db),
		IndexSubject:        q.IndexSubject.clone(db),
		Member:              q.Member.clone(db),
		Notification:        q.Notification.clone(db),
		NotificationField:   q.NotificationField.clone(db),
		OAuthClient:         q.OAuthClient.clone(db),
		Person:              q.Person.clone(db),
		PersonComment:       q.PersonComment.clone(db),
		PersonField:         q.PersonField.clone(db),
		PersonSubjects:      q.PersonSubjects.clone(db),
		PrivateMessage:      q.PrivateMessage.clone(db),
		RevisionHistory:     q.RevisionHistory.clone(db),
		RevisionText:        q.RevisionText.clone(db),
		Subject:             q.Subject.clone(db),
		SubjectCollection:   q.SubjectCollection.clone(db),
		SubjectField:        q.SubjectField.clone(db),
		SubjectRelation:     q.SubjectRelation.clone(db),
		SubjectRevision:     q.SubjectRevision.clone(db),
		SubjectTopic:        q.SubjectTopic.clone(db),
		SubjectTopicComment: q.SubjectTopicComment.clone(db),
		TimeLine:            q.TimeLine.clone(db),
		UserGroup:           q.UserGroup.clone(db),
		WebSession:          q.WebSession.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                  db,
		AccessToken:         q.AccessToken.replaceDB(db),
		App:                 q.App.replaceDB(db),
		Cast:                q.Cast.replaceDB(db),
		Character:           q.Character.replaceDB(db),
		CharacterComment:    q.CharacterComment.replaceDB(db),
		CharacterSubjects:   q.CharacterSubjects.replaceDB(db),
		EpCollection:        q.EpCollection.replaceDB(db),
		Episode:             q.Episode.replaceDB(db),
		EpisodeComment:      q.EpisodeComment.replaceDB(db),
		Friend:              q.Friend.replaceDB(db),
		Group:               q.Group.replaceDB(db),
		GroupMember:         q.GroupMember.replaceDB(db),
		GroupTopic:          q.GroupTopic.replaceDB(db),
		GroupTopicComment:   q.GroupTopicComment.replaceDB(db),
		Index:               q.Index.replaceDB(db),
		IndexComment:        q.IndexComment.replaceDB(db),
		IndexSubject:        q.IndexSubject.replaceDB(db),
		Member:              q.Member.replaceDB(db),
		Notification:        q.Notification.replaceDB(db),
		NotificationField:   q.NotificationField.replaceDB(db),
		OAuthClient:         q.OAuthClient.replaceDB(db),
		Person:              q.Person.replaceDB(db),
		PersonComment:       q.PersonComment.replaceDB(db),
		PersonField:         q.PersonField.replaceDB(db),
		PersonSubjects:      q.PersonSubjects.replaceDB(db),
		PrivateMessage:      q.PrivateMessage.replaceDB(db),
		RevisionHistory:     q.RevisionHistory.replaceDB(db),
		RevisionText:        q.RevisionText.replaceDB(db),
		Subject:             q.Subject.replaceDB(db),
		SubjectCollection:   q.SubjectCollection.replaceDB(db),
		SubjectField:        q.SubjectField.replaceDB(db),
		SubjectRelation:     q.SubjectRelation.replaceDB(db),
		SubjectRevision:     q.SubjectRevision.replaceDB(db),
		SubjectTopic:        q.SubjectTopic.replaceDB(db),
		SubjectTopicComment: q.SubjectTopicComment.replaceDB(db),
		TimeLine:            q.TimeLine.replaceDB(db),
		UserGroup:           q.UserGroup.replaceDB(db),
		WebSession:          q.WebSession.replaceDB(db),
	}
}

type queryCtx struct {
	AccessToken         *accessTokenDo
	App                 *appDo
	Cast                *castDo
	Character           *characterDo
	CharacterComment    *characterCommentDo
	CharacterSubjects   *characterSubjectsDo
	EpCollection        *epCollectionDo
	Episode             *episodeDo
	EpisodeComment      *episodeCommentDo
	Friend              *friendDo
	Group               *groupDo
	GroupMember         *groupMemberDo
	GroupTopic          *groupTopicDo
	GroupTopicComment   *groupTopicCommentDo
	Index               *indexDo
	IndexComment        *indexCommentDo
	IndexSubject        *indexSubjectDo
	Member              *memberDo
	Notification        *notificationDo
	NotificationField   *notificationFieldDo
	OAuthClient         *oAuthClientDo
	Person              *personDo
	PersonComment       *personCommentDo
	PersonField         *personFieldDo
	PersonSubjects      *personSubjectsDo
	PrivateMessage      *privateMessageDo
	RevisionHistory     *revisionHistoryDo
	RevisionText        *revisionTextDo
	Subject             *subjectDo
	SubjectCollection   *subjectCollectionDo
	SubjectField        *subjectFieldDo
	SubjectRelation     *subjectRelationDo
	SubjectRevision     *subjectRevisionDo
	SubjectTopic        *subjectTopicDo
	SubjectTopicComment *subjectTopicCommentDo
	TimeLine            *timeLineDo
	UserGroup           *userGroupDo
	WebSession          *webSessionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		AccessToken:         q.AccessToken.WithContext(ctx),
		App:                 q.App.WithContext(ctx),
		Cast:                q.Cast.WithContext(ctx),
		Character:           q.Character.WithContext(ctx),
		CharacterComment:    q.CharacterComment.WithContext(ctx),
		CharacterSubjects:   q.CharacterSubjects.WithContext(ctx),
		EpCollection:        q.EpCollection.WithContext(ctx),
		Episode:             q.Episode.WithContext(ctx),
		EpisodeComment:      q.EpisodeComment.WithContext(ctx),
		Friend:              q.Friend.WithContext(ctx),
		Group:               q.Group.WithContext(ctx),
		GroupMember:         q.GroupMember.WithContext(ctx),
		GroupTopic:          q.GroupTopic.WithContext(ctx),
		GroupTopicComment:   q.GroupTopicComment.WithContext(ctx),
		Index:               q.Index.WithContext(ctx),
		IndexComment:        q.IndexComment.WithContext(ctx),
		IndexSubject:        q.IndexSubject.WithContext(ctx),
		Member:              q.Member.WithContext(ctx),
		Notification:        q.Notification.WithContext(ctx),
		NotificationField:   q.NotificationField.WithContext(ctx),
		OAuthClient:         q.OAuthClient.WithContext(ctx),
		Person:              q.Person.WithContext(ctx),
		PersonComment:       q.PersonComment.WithContext(ctx),
		PersonField:         q.PersonField.WithContext(ctx),
		PersonSubjects:      q.PersonSubjects.WithContext(ctx),
		PrivateMessage:      q.PrivateMessage.WithContext(ctx),
		RevisionHistory:     q.RevisionHistory.WithContext(ctx),
		RevisionText:        q.RevisionText.WithContext(ctx),
		Subject:             q.Subject.WithContext(ctx),
		SubjectCollection:   q.SubjectCollection.WithContext(ctx),
		SubjectField:        q.SubjectField.WithContext(ctx),
		SubjectRelation:     q.SubjectRelation.WithContext(ctx),
		SubjectRevision:     q.SubjectRevision.WithContext(ctx),
		SubjectTopic:        q.SubjectTopic.WithContext(ctx),
		SubjectTopicComment: q.SubjectTopicComment.WithContext(ctx),
		TimeLine:            q.TimeLine.WithContext(ctx),
		UserGroup:           q.UserGroup.WithContext(ctx),
		WebSession:          q.WebSession.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
