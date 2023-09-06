// Code generated by goctl. DO NOT EDIT.
package types

type UserRegisterRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserRegisterResponse struct {
	RespStatus
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserLoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type UserLoginResponse struct {
	RespStatus
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token"`
}

type UserResponse struct {
	RespStatus
	User User `json:"user"`
}

type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count,optional"`
	FollowerCount   int64  `json:"follower_count,optional"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar,optional"`
	BackgroundImage string `json:"background_image,optional"`
	Signature       string `json:"signature,optional"`
	TotalFavorited  int64  `json:"total_favorited,optional"`
	WorkCount       int64  `json:"work_count,optional"`
	FavoriteCount   int64  `json:"favorite_count,optional"`
}

type Video struct {
	ID            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayURL       string `json:"play_url"`
	CoverURL      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type RespStatus struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,optional"`
}

type FeedRequest struct {
	LatestTime int64  `form:"latest_time,optional"`
	Token      string `form:"token,optional"`
}

type FeedResponse struct {
	RespStatus
	VideoList []Video `json:"video_list"`
	NextTime  int64   `json:"next_time,optional"`
}

type PublishActionRequest struct {
	Token string `form:"token"`
	Title string `form:"title"`
}

type PublishActionResponse struct {
	RespStatus
}

type PublishListRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token,optional"`
}

type PublishListResponse struct {
	RespStatus
	VideoList []Video `json:"video_list"`
}

type FavoriteActionRequest struct {
	Token      string `form:"token"`
	VideoID    int64  `form:"video_id,range=[0:]"`
	ActionType int32  `form:"action_type,options=[1,2]"`
}

type FavoriteActionResponse struct {
	RespStatus
}

type FavoriteListRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token,optional"`
}

type FavoriteListResponse struct {
	RespStatus
	VideoList []Video `json:"video_list"`
}

type CommentActionRequest struct {
	Token       string `form:"token"`
	VideoID     int64  `form:"video_id,range=[0:]"`
	ActionType  int32  `form:"action_type,options=[1,2]"`
	CommentText string `form:"comment_text,optional"`
	CommentID   int64  `form:"comment_id,range=[0:],optional"`
}

type CommentActionResponse struct {
	RespStatus
	Comment Comment `json:"comment,optional"`
}

type Comment struct {
	ID         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentListRequest struct {
	Token   string `form:"token,optional"`
	VideoID int64  `form:"video_id,range=[0:]"`
}

type CommentListResponse struct {
	RespStatus
	CommentList []Comment `json:"comment_list"`
}

type RelationActionRequest struct {
	Token      string `form:"token"`
	ToUserID   int64  `form:"to_user_id,range=[0:]"`
	ActionType int32  `form:"action_type,options=[1,2]"`
}

type RelationActionResponse struct {
	RespStatus
}

type RelationFollowListRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token,optional"`
}

type RelationFollowListResponse struct {
	RespStatus
	UserList []User `json:"user_list"`
}

type RelationFollowerListRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token,optional"`
}

type RelationFollowerListResponse struct {
	RespStatus
	UserList []User `json:"user_list"`
}

type RelationFriendListRequest struct {
	UserID int64  `form:"user_id,range=[0:]"`
	Token  string `form:"token"`
}

type RelationFriendListResponse struct {
	RespStatus
	UserList []FriendUser `json:"user_list"`
}

type FriendUser struct {
	User
	Message string `json:"message,optional"`
	MsgType int64  `json:"msgType"`
}

type MessageChatRequest struct {
	Token      string `form:"token"`
	ToUserID   int64  `form:"to_user_id,range=[0:]"`
	PreMsgTime int64  `form:"pre_msg_time"`
}

type MessageChatResponse struct {
	RespStatus
	MessageList []Message `json:"message_list"`
}

type Message struct {
	ID         int64  `json:"id"`
	ToUserID   int64  `json:"to_user_id"`
	FromUserID int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time,optional"`
}

type MessageActionRequest struct {
	Token      string `form:"token"`
	ToUserID   int64  `form:"to_user_id,range=[0:]"`
	ActionType int32  `form:"action_type,options=[1]"`
	Content    string `form:"content"`
}

type MessageActionResponse struct {
	RespStatus
}
