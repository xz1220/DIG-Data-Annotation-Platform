package model


type Points struct {
	Order int     `json:"order"` //注意这里可能出现错误，因为会出现关键字冲突
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

// Data 接受前端参数的结构体
type DataForRequest struct {
	Point []Points `json:"point"`

	LabelID   int64  `json:"labelId"`
	LabelType int64  `json:"labelType"`
	DataDesc  string `json:"dataDesc"`
	IScrowd   int    `json:"iscrowd"`
}

type LabelData struct {
	Data []*DataForRequest `json:"data"`

	ImageIDString string `json:"imageId"`
	UserID        int64  `json:"userId"`
}

type Data struct {
	DataID    int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:data_id" form:"user_id"`
	ImageID   int    `gorm:"column:image_id" form:"image_id"`
	LabelID   int    `gorm:"column:label_id" form:"label_id"`
	UserID    int    `gorm:"column:user_id" form:"user_id"`
	DataDesc  string `gorm:"type:varchar(1024);column:data_desc" form:"data_desc"`
	LabelType int    `gorm:"column:label_type" form:"label_type"`
	Iscrowd   int    `gorm:"iscrowd" form:"iscrowd"`

	Order int     `gorm:"column:order" form:"order"` //注意这里可能出现错误，因为会出现关键字冲突
	X     float64 `gorm:"column:x" form:"x"`
	Y     float64 `gorm:"column:y" form:"y"`
}

type DataForResponse struct {
	DataID    int64  `json:"dataId"`
	ImageID   int    `json:"imageId"`
	LabelID   int    `json:"labelId"`
	UserID    int    `json:"userId"`
	DataDesc  string `json:"dataDesc"`
	LabelType int    `json:"labelType"`
	Iscrowd   int    `json:"iscrowd"`

	Point []Points `json:"point"`
}

type CocoInfo struct {
	Year        int64  `json:"year"`
	DataCreated string `json:"data_created"`
}

type CocoImage struct {
	ID       int64  `json:"id"`
	FileName string `json:"file_name"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
}

type CocoAnnotation struct {
	ID           int64     `json:"id"`
	ImageID      int64     `json:"image_id"`
	CategoryID   int64     `json:"category_id"`
	Area         float64   `json:"area"`
	Iscrowd      int       `json:"iscrowd"`
	Segmentation []float64 `json:"segmentation"`
	BBox         []float64 `json:"bbox"`
	Desc         string    `json:"desc"`
}

type CocoCategory struct {
	SuperCategory string `json:"supercategory"`
	ID            int64  `json:"id"`
	Name          string `json:"name"`
}

type CocoDataSet struct {
	Info        CocoInfo         `json:"info"`
	Images      []CocoImage      `json:"images"`
	Annotations []CocoAnnotation `json:"annotations"`
	Categories  []CocoCategory   `json:"categories"`

	Keypoints []string `json:"keypoints"`
}



type PageData struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type TaskResponse struct {
	Users    []*User `json:"users"`
	Reviewer []*User `json:"reviewers"`

	TaskID      int64   `json:"taskId"`
	TaskName    string  `json:"taskName"`
	TaskDesc    string  `json:"taskDesc"`
	ImageNumber int64   `json:"imageNumber"`
	TaskType    int64   `json:"taskType"`
	Finish      int64   `json:"finish"`
	UserIDs     []int64 `json:"userIds"`
	ReviewerIDs []int64 `json:"reviewerIds"`
	LabelIDs    []int64 `json:"labelIds"`
}

type TaskRequest struct {
	TaskID      int64   `json:"taskId"`
	TaskName    string  `json:"taskName"`
	TaskDesc    string  `json:"taskDesc"`
	UserIDs     []int64 `json:"userIds"`
	ReviewerIDs []int64 `json:"reviewerIds"`
	LabelIDs    []int64 `json:"labelIds"`
}

