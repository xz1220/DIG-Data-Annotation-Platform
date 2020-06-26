package model

// Image is tabel storing the information of ids, including usercomfirmID and taskID
type Image struct {
	ImageID       int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:image_id" form:"image_id" json:"imageId"`
	ImageName     string `gorm:"type:varchar(1024);column:image_name" form:"image_name" json:"imageName"`
	ImageThumb    string `gorm:"type:varchar(1024);column:image_thumb" form:"image_thumb"  json:"imageThumb"`
	UserComfirmID int64  `gorm:"column:user_confirm_id" form:"user_confirm_id"  json:"userConfirmId"`
	TaskID        int64  `gorm:"column:task_id" form:"task_id"  json:"taskId"`
	Width         int64  `gorm:"column:width" form:"width"  json:"width"`
	Height        int64  `gorm:"column:height" form:"height"  json:"height"`
}

//TableName rename the table.
func (Image) TableName() string {
	return "image"
}

// ImageData is
type ImageData struct {
	DataID    int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:data_id" form:"user_id"`
	ImageID   int64  `gorm:"column:image_id" form:"image_id"`
	LabelID   int64  `gorm:"column:label_id" form:"label_id"`
	UserID    int64  `gorm:"column:user_id" form:"user_id"`
	DataDesc  string `gorm:"type:varchar(1024);column:data_desc" form:"data_desc"`
	LabelType int64  `gorm:"column:label_type" form:"label_type"`
	Iscrowd   int    `gorm:"iscrowd" form:"iscrowd"`
}

//
func (ImageData) TableName() string {
	return "imagedata" //注意大小写
}

// ImageDataPoints is ..
type ImageDataPoints struct {
	DataID  int64   `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:data_id" form:"data_id"`
	Order   int     `gorm:"column:order" form:"order"` //注意这里可能出现错误，因为会出现关键字冲突
	X       float64 `gorm:"column:x" form:"x"`
	Y       float64 `gorm:"column:y" form:"y"`
	ImageID int64   `gorm:"column:image_id" form:"image_id"`
	UserID  int64   `gorm:"column:user_id" form:"user_id"`
}

func (ImageDataPoints) TableName() string {
	return "imagedatapoints"
}

type Imagedatarle struct {
	DataID  int64  `gorm:"column:data_id" form:"data_id"`
	ImageID int64  `gorm:"column:image_id" form:"image_id"`
	UserID  int64  `gorm:"column:user_id" form:"user_id"`
	DataRle string `gorm:"type:longtext;column:data_id" form:"data_rle"`
}

func (Imagedatarle) TableName() string {
	return "imagedatarle"
}

type RleDatas struct {
	Counts []int64
	Size   []int64
}

type TempRleData struct {
	DataID int64
	Data   string
}

type Imagelabel struct {
	LabelID    int64  `gorm:"primary_key;AUTO_INCREMENT;unique_index;column:label_id" form:"label_id" json:"labelId"`
	LabelName  string `gorm:"type:varchar(50);column:label_name" form:"label_name" json:"labelName"`
	LabelType  int    `gorm:"column:label_type" form:"label_type" json:"labelType"`
	LabelColor string `gorm:"type:varchar(20);column:label_color" form:"label_color" json:"labelColor"`
}

func (Imagelabel) TableName() string {
	return "imagelabel"
}

type Points struct {
	Order int     `json:"order"` //注意这里可能出现错误，因为会出现关键字冲突
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

// Data 接受前端参数的结构体
type DataForRequest struct {
	LabelID   int64    `json:"labelId"`
	LabelType int64    `json:"labelType"`
	DataDesc  string   `json:"dataDesc"`
	IScrowd   int      `json:"iscrowd"`
	Point     []Points `json:"point"`
}

type LabelData struct {
	ImageIDString string            `json:"imageId"`
	Data          []*DataForRequest `json:"data"`
	UserID        int64             `json:"userId"`
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
	Segmentation string    `json:"segmentation"`
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
}
