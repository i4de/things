package relationDB

import (
	"database/sql"
	"gitee.com/i-Things/share/stores"
	"github.com/i-Things/things/service/dmsvr/internal/domain/productCustom"
	"time"
)

type DmExample struct {
	ID int64 `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"` // id编号
}

// 设备信息表
type DmDeviceInfo struct {
	ID             int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	TenantCode     stores.TenantCode `gorm:"column:tenant_code;index;type:VARCHAR(50);NOT NULL"`                              // 租户编码
	ProjectID      stores.ProjectID  `gorm:"column:project_id;index:project_id_area_id;type:bigint;default:0;NOT NULL"`       // 项目ID(雪花ID)
	AreaID         stores.AreaID     `gorm:"column:area_id;index:project_id_area_id;type:bigint;default:0;NOT NULL"`          // 项目区域ID(雪花ID)
	ProductID      string            `gorm:"column:product_id;type:char(11);uniqueIndex:product_id_deviceName;NOT NULL"`      // 产品id
	DeviceName     string            `gorm:"column:device_name;uniqueIndex:product_id_deviceName;type:varchar(100);NOT NULL"` // 设备名称
	DeviceAlias    string            `gorm:"column:device_alias;type:varchar(100);NOT NULL"`                                  // 设备别名
	Position       stores.Point      `gorm:"column:position;type:point;NOT NULL"`                                             // 设备的位置(默认百度坐标系BD09)
	Secret         string            `gorm:"column:secret;type:varchar(50);NOT NULL"`                                         // 设备秘钥
	Cert           string            `gorm:"column:cert;type:varchar(512);NOT NULL"`                                          // 设备证书
	Imei           string            `gorm:"column:imei;type:varchar(15);NOT NULL"`                                           // IMEI号信息
	Mac            string            `gorm:"column:mac;type:varchar(17);NOT NULL"`                                            // MAC号信息
	Version        string            `gorm:"column:version;type:varchar(64);NOT NULL"`                                        // 固件版本
	HardInfo       string            `gorm:"column:hard_info;type:varchar(64);NOT NULL"`                                      // 模组硬件型号
	SoftInfo       string            `gorm:"column:soft_info;type:varchar(64);NOT NULL"`                                      // 模组软件版本
	MobileOperator int64             `gorm:"column:mobile_operator;type:smallint;default:1;NOT NULL"`                         // 移动运营商:1)移动 2)联通 3)电信 4)广电
	Phone          sql.NullString    `gorm:"column:phone;type:varchar(20)"`                                                   // 手机号
	Iccid          sql.NullString    `gorm:"column:iccid;type:varchar(20)"`                                                   // SIM卡卡号
	Address        string            `gorm:"column:address;type:varchar(512);NOT NULL"`                                       // 所在地址
	Tags           map[string]string `gorm:"column:tags;type:json;serializer:json;NOT NULL;default:'{}'"`                     // 设备标签
	IsOnline       int64             `gorm:"column:is_online;type:smallint;default:2;NOT NULL"`                               // 是否在线,1是2否
	FirstLogin     sql.NullTime      `gorm:"column:first_login"`                                                              // 激活时间
	LastLogin      sql.NullTime      `gorm:"column:last_login"`                                                               // 最后上线时间
	LogLevel       int64             `gorm:"column:log_level;type:smallint;default:1;NOT NULL"`                               // 日志级别:1)关闭 2)错误 3)告警 4)信息 5)调试
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:product_id_deviceName"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"` // 添加外键
}

func (m *DmDeviceInfo) TableName() string {
	return "dm_device_info"
}

// 产品信息表
type DmProductInfo struct {
	ProductID    string            `gorm:"column:product_id;type:char(11);primary_key;NOT NULL"`          // 产品id
	ProductName  string            `gorm:"column:product_name;uniqueIndex:pn;type:varchar(100);NOT NULL"` // 产品名称
	ProductImg   string            `gorm:"column:product_img;type:varchar(200)"`                          // 产品图片
	ProductType  int64             `gorm:"column:product_type;type:smallint;default:1"`                   // 产品状态:1:开发中,2:审核中,3:已发布
	AuthMode     int64             `gorm:"column:auth_mode;type:smallint;default:1"`                      // 认证方式:1:账密认证,2:秘钥认证
	DeviceType   int64             `gorm:"column:device_type;index;type:smallint;default:1"`              // 设备类型:1:设备,2:网关,3:子设备
	CategoryID   int64             `gorm:"column:category_id;type:integer;default:1"`                     // 产品品类
	NetType      int64             `gorm:"column:net_type;type:smallint;default:1"`                       // 通讯方式:1:其他,2:wi-fi,3:2G/3G/4G,4:5G,5:BLE,6:LoRaWAN
	DataProto    int64             `gorm:"column:data_proto;type:smallint;default:1"`                     // 数据协议:1:自定义,2:数据模板
	ProtocolID   int64             `gorm:"column:protocol_id;type:bigint;default:1"`                      // 协议名称: 如果为空则为iThings标准协议
	AutoRegister int64             `gorm:"column:auto_register;type:smallint;default:1"`                  // 动态注册:1:关闭,2:打开,3:打开并自动创建设备
	Secret       string            `gorm:"column:secret;type:varchar(50)"`                                // 动态注册产品秘钥
	Desc         string            `gorm:"column:description;type:varchar(200)"`                          // 描述
	DevStatus    string            `gorm:"column:dev_status;type:varchar(20);NOT NULL"`                   // 产品状态
	Tags         map[string]string `gorm:"column:tags;type:json;serializer:json;NOT NULL;default:'{}'"`   // 产品标签
	stores.NoDelTime
	DeletedTime  stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pn"`
	ProtocolInfo *DmProtocolInfo    `gorm:"foreignKey:ID;references:ProtocolID"` // 添加外键
	Category     *DmProductCategory `gorm:"foreignKey:ID;references:CategoryID"` // 添加外键
	//Devices []*DmDeviceInfo    `gorm:"foreignKey:ProductID;references:ProductID"` // 添加外键

}

func (m *DmProductInfo) TableName() string {
	return "dm_product_info"
}

type DmProductCategory struct {
	ID      int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Name    string `gorm:"column:name;uniqueIndex:pn;type:varchar(100);NOT NULL"` // 产品品类名称
	Desc    string `gorm:"column:desc;type:varchar(200)"`                         // 描述
	HeadImg string `gorm:"column:head_img;type:varchar(200)"`                     // 图片
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pn"`
}

func (m *DmProductCategory) TableName() string {
	return "dm_product_category"
}

// 自定义协议表
type DmProtocolInfo struct {
	ID           int64    `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Name         string   `gorm:"column:name;uniqueIndex:pn;type:varchar(100);NOT NULL"`            // 协议名称
	Protocol     string   `gorm:"column:protocol;type:varchar(100)"`                                // 协议: mqtt,tcp,udp
	ProtocolType string   `gorm:"column:protocol_type;type:varchar(100);default:iThings"`           // 协议类型: iThings,iThings-thingsboard
	Desc         string   `gorm:"column:desc;type:varchar(200)"`                                    // 描述
	Endpoints    []string `gorm:"column:endpoints;type:json;serializer:json;NOT NULL;default:'[]'"` // 协议端点,如果填写了优先使用该字段
	EtcdKey      string   `gorm:"column:etcd_key;type:varchar(200);default:null"`                   //服务etcd发现的key etcd key
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pn"`
}

func (m *DmProtocolInfo) TableName() string {
	return "dm_protocol_info"
}

// 产品自定义协议表
type DmProductCustom struct {
	ID              int64                        `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID       string                       `gorm:"column:product_id;uniqueIndex:pn;type:char(11);NOT NULL"` // 产品id
	ScriptLang      int64                        `gorm:"column:script_lang;type:smallint;default:1"`              // 脚本语言类型 1:JavaScript 2:lua 3:python
	CustomTopics    []*productCustom.CustomTopic `gorm:"column:custom_topics;type:json"`                          // 自定义topic数组
	TransformScript string                       `gorm:"column:transform_script;type:text"`                       // 协议转换脚本
	LoginAuthScript string                       `gorm:"column:login_auth_script;type:text"`                      // 登录认证脚本
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pn"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (m *DmProductCustom) TableName() string {
	return "dm_product_custom"
}

// 产品物模型表
type DmProductSchema struct {
	ID        int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID string `gorm:"column:product_id;uniqueIndex:identifier;index:product_id_type;type:char(11);NOT NULL"` // 产品id
	Tag       int64  `gorm:"column:tag;type:smallint;default:1"`                                                    // 物模型标签 1:自定义 2:可选 3:必选  必选不可删除
	DmSchemaCore
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:identifier"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (m *DmProductSchema) TableName() string {
	return "dm_product_schema"
}

type DmSchemaCore struct {
	Type              int64  `gorm:"column:type;index:product_id_type;type:smallint;default:1"`           // 物模型类型 1:property属性 2:event事件 3:action行为
	Identifier        string `gorm:"column:identifier;uniqueIndex:identifier;type:varchar(100);NOT NULL"` // 标识符
	ExtendConfig      string `gorm:"column:extend_config;type:json;default:'{}'"`                         //拓展参数
	Required          int64  `gorm:"column:required;type:smallint;default:2"`                             // 是否必须,1是 2否
	Name              string `gorm:"column:name;type:varchar(100);NOT NULL"`                              // 功能名称
	Desc              string `gorm:"column:desc;type:varchar(200)"`                                       // 描述
	IsCanSceneLinkage int64  `gorm:"column:is_can_scene_linkage;type:smallint;default:1"`                 // 是否放到场景联动中
	IsShareAuthPerm   int64  `gorm:"column:is_share_auth_perm;type:smallint;default:1"`                   // 分享是否需要校验权限
	IsHistory         int64  `gorm:"column:is_history;type:smallint;default:1"`                           // 是否存储历史记录
	Affordance        string `gorm:"column:affordance;type:json;NOT NULL"`                                // 各类型的自定义功能定义
	Order             int64  `gorm:"column:order;type:BIGINT;default:1;NOT NULL"`                         // 左侧table排序序号
}

// 通用物模型表
type DmCommonSchema struct {
	ID int64 `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	DmSchemaCore
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:identifier"`
}

func (m *DmCommonSchema) TableName() string {
	return "dm_common_schema"
}

// 设备分组信息表
type DmGroupInfo struct {
	ID         int64             `gorm:"column:id;primary_key;AUTO_INCREMENT;type:bigint"`                               // 分组ID
	TenantCode stores.TenantCode `gorm:"column:tenant_code;uniqueIndex:tc_ac;default:default;type:VARCHAR(50);NOT NULL"` // 租户编码
	ProjectID  stores.ProjectID  `gorm:"column:project_id;uniqueIndex:tc_ac;type:bigint;default:2;NOT NULL"`             // 项目ID(雪花ID)
	AreaID     stores.AreaID     `gorm:"column:area_id;uniqueIndex:tc_ac;type:bigint;default:2;NOT NULL"`                // 项目区域ID(雪花ID)
	ParentID   int64             `gorm:"column:parent_id;type:bigint;default:0;NOT NULL"`                                // 父组ID 0-根组
	ProductID  string            `gorm:"column:product_id;type:char(11);default:'';NOT NULL"`                            // 产品id,为空则不限定分组内的产品类型
	Name       string            `gorm:"column:name;uniqueIndex:tc_ac;default:'';type:varchar(100);NOT NULL"`            // 分组名称
	Desc       string            `gorm:"column:desc;type:varchar(200);default:''"`                                       // 描述
	Tags       map[string]string `gorm:"column:tags;type:json;serializer:json;NOT NULL;default:'{}'"`                    // 分组标签
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:tc_ac"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (m *DmGroupInfo) TableName() string {
	return "dm_group_info"
}

// 分组与设备关系表
type DmGroupDevice struct {
	ID         int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	TenantCode stores.TenantCode `gorm:"column:tenant_code;index;type:VARCHAR(50);NOT NULL"`                                        // 租户编码
	ProjectID  stores.ProjectID  `gorm:"column:project_id;index;type:bigint;default:0;NOT NULL"`                                    // 项目ID(雪花ID)
	AreaID     stores.AreaID     `gorm:"column:area_id;index:project_id_area_id;type:bigint;default:0;NOT NULL"`                    // 项目区域ID(雪花ID)
	GroupID    int64             `gorm:"column:group_id;uniqueIndex:group_id_product_id_device_name;type:bigint;NOT NULL"`          // 分组ID
	ProductID  string            `gorm:"column:product_id;uniqueIndex:group_id_product_id_device_name;type:char(11);NOT NULL"`      // 产品id
	DeviceName string            `gorm:"column:device_name;uniqueIndex:group_id_product_id_device_name;type:varchar(100);NOT NULL"` // 设备名称
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:group_id_product_id_device_name"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"`
	Device      *DmDeviceInfo      `gorm:"foreignKey:ProductID,DeviceName;references:ProductID,DeviceName"`
}

func (m *DmGroupDevice) TableName() string {
	return "dm_group_device"
}

// 网关与子设备关系表
type DmGatewayDevice struct {
	TenantCode        stores.TenantCode `gorm:"column:tenant_code;index;type:VARCHAR(50);NOT NULL"` // 租户编码
	ID                int64             `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	GatewayProductID  string            `gorm:"column:gateway_product_id;uniqueIndex:gpi_pdn_pi_dn;type:char(11);NOT NULL"`      // 网关产品id
	GatewayDeviceName string            `gorm:"column:gateway_device_name;uniqueIndex:gpi_pdn_pi_dn;type:varchar(100);NOT NULL"` // 网关设备名称
	ProductID         string            `gorm:"column:product_id;uniqueIndex:gpi_pdn_pi_dn;type:char(11);NOT NULL"`              // 子设备产品id
	DeviceName        string            `gorm:"column:device_name;uniqueIndex:gpi_pdn_pi_dn;type:varchar(100);NOT NULL"`         // 子设备名称
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:gpi_pdn_pi_dn"`
	Device      *DmDeviceInfo      `gorm:"foreignKey:ProductID,DeviceName;references:ProductID,DeviceName"`
	Gateway     *DmDeviceInfo      `gorm:"foreignKey:ProductID,DeviceName;references:GatewayProductID,GatewayDeviceName"`
}

func (m *DmGatewayDevice) TableName() string {
	return "dm_gateway_device"
}

// 产品远程配置表
type DmProductRemoteConfig struct {
	ID        int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID string `gorm:"column:product_id;uniqueIndex:pn;type:char(11);NOT NULL"` // 产品id
	Content   string `gorm:"column:content;type:json;NOT NULL"`                       // 配置内容
	stores.NoDelTime
	DeletedTime stores.DeletedTime `gorm:"column:deleted_time;uniqueIndex:pn"`
	ProductInfo *DmProductInfo     `gorm:"foreignKey:ProductID;references:ProductID"`
}

func (m *DmProductRemoteConfig) TableName() string {
	return "dm_product_remote_config"
}

// 升级任务表
type DmOtaTask struct {
	ID          int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID   string `gorm:"column:product_id;type:char(11);NOT NULL"`             // 产品id
	FirmwareID  int64  `gorm:"column:firmware_id;type:bigint;NOT NULL"`              // 固件id
	TaskUid     string `gorm:"column:task_uid;type:varchar(64)"`                     // 任务编号
	Type        int64  `gorm:"column:type;type:smallint;default:1;NOT NULL"`         // 升级范围1全部设备2定向升级
	UpgradeType int64  `gorm:"column:upgrade_type;type:smallint;default:1;NOT NULL"` // 升级策略:1静态升级2动态升级
	AutoRepeat  int64  `gorm:"column:auto_repeat;type:smallint;default:1;NOT NULL"`  // 是否自动重试,1:不,2自动重试
	Status      int64  `gorm:"column:status;type:smallint;default:1;NOT NULL"`       // 升级状态:1未升级2升级中3完成4已取消
	DeviceList  string `gorm:"column:device_list;type:json;NOT NULL"`                // 指定升级设备
	VersionList string `gorm:"column:version_list;type:json;NOT NULL"`               // 指定待升级版本
	stores.Time
}

func (m *DmOtaTask) TableName() string {
	return "dm_ota_task"
}

// 升级包附件列表
type DmOtaFirmwareFile struct {
	ID         int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	Name       string `gorm:"column:name;type:varchar(64)"`                // 附件名称
	FirmwareID int64  `gorm:"column:firmware_id;type:bigint;NOT NULL"`     // 固件id
	Size       int64  `gorm:"column:size;type:bigint;NOT NULL"`            // 文件大小单位bit
	Storage    string `gorm:"column:storage;type:varchar(15);NOT NULL"`    // 存储平台:minio/aliyun
	Host       string `gorm:"column:host;type:varchar(100);NOT NULL"`      // host
	FilePath   string `gorm:"column:file_path;type:varchar(100);NOT NULL"` // 文件路径,拿来下载文件
	Signature  string `gorm:"column:signature;type:char(32);NOT NULL"`     // 签名值
	stores.Time
}

func (m *DmOtaFirmwareFile) TableName() string {
	return "dm_ota_firmware_file"
}

// ota升级记录
type DmOtaTaskDevices struct {
	ID            int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	FirmwareID    int64  `gorm:"column:firmware_id;type:bigint;NOT NULL"`             // 固件id
	TaskUid       string `gorm:"column:task_uid;type:varchar(64);NOT NULL"`           // 任务批次
	ProductID     string `gorm:"column:product_id;type:char(11);NOT NULL"`            // 产品id
	DeviceName    string `gorm:"column:device_name;type:varchar(100);NOT NULL"`       // 设备编号
	Version       string `gorm:"column:version;type:varchar(64)"`                     // 当前版本
	TargetVersion string `gorm:"column:target_version;type:varchar(64);NOT NULL"`     // 升级包的版本
	Status        int64  `gorm:"column:status;type:integer;default:101"`              // 升级状态:101待确认 201/202/203待推送 301已推送 401升级中 501升级成功 601升级失败 701已取消
	RetryCount    int64  `gorm:"column:retry_count;type:smallint;default:0;NOT NULL"` // 重试次数,计划最多20次
	Step          int64  `gorm:"column:step;type:integer;default:0;NOT NULL"`         // OTA升级进度。1~100的整数升级进度百分比,-1升级失败,-2下载失败,-3校验失败,-4烧写失败
	Desc          string `gorm:"column:desc;type:varchar(200)"`                       // 状态详情
	stores.Time
}

func (m *DmOtaTaskDevices) TableName() string {
	return "dm_ota_task_devices"
}

// 产品固件升级包信息表
type DmOtaFirmware struct {
	ID         int64  `gorm:"column:id;type:bigint;primary_key;AUTO_INCREMENT"`
	ProductID  string `gorm:"column:product_id;type:char(11);NOT NULL"`        // 产品id
	Version    string `gorm:"column:version;type:varchar(64)"`                 // 固件版本
	SrcVersion string `gorm:"column:src_version;type:varchar(64)"`             // 待升级版本号
	Module     string `gorm:"column:module;type:varchar(64)"`                  // 模块名称
	Name       string `gorm:"column:name;type:varchar(64)"`                    // 固件名称
	Desc       string `gorm:"column:desc;type:varchar(200)"`                   // 描述
	Status     int64  `gorm:"column:status;type:bigint;NOT NULL"`              //升级包状态，-1：不需要验证，0：未验证，1：已验证，2：验证中，3：验证失败
	TotalSize  int64  `gorm:"column:total_size;type:bigint;NOT NULL"`          // 升级包总大小
	IsDiff     int64  `gorm:"column:is_diff;type:smallint;default:1;NOT NULL"` // 是否差分包,1:整包,2:差分
	SignMethod string `gorm:"column:sign_method;type:varchar(20);NOT NULL"`    // 签名方式:MD5/SHA256
	Extra      string `gorm:"column:extra;type:varchar(256)"`                  // 自定义推送参数
	stores.Time
}

func (m *DmOtaFirmware) TableName() string {
	return "dm_ota_firmware"
}

// 设备影子表
type DmDeviceShadow struct {
	ID                int64      `gorm:"column:id;type:BIGINT;primary_key;AUTO_INCREMENT"`
	ProductID         string     `gorm:"column:product_id;uniqueIndex:pi_dn_di;type:CHAR(11);NOT NULL"`
	DeviceName        string     `gorm:"column:device_name;uniqueIndex:pi_dn_di;type:VARCHAR(100);NOT NULL"`
	DataID            string     `gorm:"column:data_id;uniqueIndex:pi_dn_di;type:VARCHAR(100);NOT NULL"`
	Value             string     `gorm:"column:value;type:VARCHAR(100);default:NULL"`
	UpdatedDeviceTime *time.Time `gorm:"column:updated_device_time;default:NULL"`
	stores.Time
}

func (m *DmDeviceShadow) TableName() string {
	return "dm_device_shadow"
}

// DMOTAjob 表示OTA升级任务的信息
type DmOtaJob struct {
	ID                 int64  `gorm:"column:id;type:BIGINT;primary_key;AUTO_INCREMENT"`
	FirmwareId         int64  `gorm:"column:firmware_id"`          // 升级包ID，升级包的唯一标识符。
	ProductId          string `gorm:"column:product_id"`           // 升级包所属产品的ProductId。
	JobType            int    `gorm:"column:job_type"`             // 升级包所属产品的JobType。
	JobStatus          int    `gorm:"column:job_status"`           // 批次状态,PLANNED：计划中,IN_PROGRESS,执行中，COMPLETED，已完成，CANCELED，已经取消
	UpgradeType        int    `gorm:"column:upgrade_type"`         // 升级策略，0-静态，1-动态
	TargetSelection    string `gorm:"column:target_selection"`     // 升级范围。可选值：0-ALL、1-SPECIFIC、2-GRAY、3-GROUP。
	SrcVersion         string `gorm:"column:src_version"`          // 待升级版本号列表。最多可传入10个版本号。用逗号分隔多个版本号
	ScheduleTime       int64  `gorm:"column:schedule_time"`        // 指定发起OTA升级的时间，单位为毫秒。
	RetryInterval      int    `gorm:"column:retry_interval"`       // 设备升级失败后，自动重试的时间间隔，单位为分钟。
	RetryCount         int    `gorm:"column:retry_count"`          // 自动重试次数。1/2/5次
	TimeoutInMinutes   int    `gorm:"column:timeout_in_minutes"`   // 设备升级超时时间，单位为分钟。
	MaximumPerMinute   int    `gorm:"column:maximum_per_minute"`   // 每分钟最多向多少个设备推送升级包下载URL。
	GrayPercent        string `gorm:"column:gray_percent"`         // 灰度比例。设置灰度升级时有效。
	TargetDeviceName   string `gorm:"column:target_device_name"`   // 定向升级的设备名称列表。最多可传入200个设备名称。以逗号分隔
	ScheduleFinishTime int64  `gorm:"column:schedule_finish_time"` // 指定结束升级的时间，单位为毫秒。
	OverwriteMode      int    `gorm:"column:overwrite_mode"`       // 是否覆盖之前的升级任务。取值：0（不覆盖）、1（覆盖）。
	DnListFileUrl      string `gorm:"column:dn_list_file_url"`     // 定向升级设备列表文件的URL。发起定向升级任务时使用。
	NeedPush           int    `gorm:"column:need_push"`            // 物联网平台是否主动向设备推送升级任务。
	NeedConfirm        int    `gorm:"column:need_confirm"`         // 是否需要App确认OTA升级。
	GroupId            string `gorm:"column:group_id"`             // 分组ID。仅发起分组升级任务时使用。
	GroupType          string `gorm:"column:group_type"`           // 分组类型，仅可取值LINK_PLATFORM。仅发起分组升级任务时使用。
	DownloadProtocol   string `gorm:"column:download_protocol"`    // 升级包下载协议。可选：0-HTTPS、1-MQTT。
	MultiModuleMode    int    `gorm:"column:multi_module_mode"`    // 设备是否支持多模块同时升级。
	stores.Time
}

func (m *DmOtaJob) TableName() string {
	return "dm_ota_job"
}

type DmOtaUpgradeTask struct {
	ID          int64     `gorm:"column:id;type:BIGINT;primary_key;AUTO_INCREMENT"`
	DestVersion string    `gorm:"column:dest_version" ` // 升级的目标OTA升级包版本
	DeviceName  string    `gorm:"column:device_name" `  // 设备名称
	FirmwareId  int64     `gorm:"column:firmware_id"`   // 升级包ID
	JobId       int64     `gorm:"column:job_id" `       // 升级批次ID
	ProductId   string    `gorm:"column:product_id" `   // 设备所属产品的ProductKey
	ProductName string    `gorm:"column:product_name"`  // 设备所属产品的名称
	Step        int64     `gorm:"column:step"`          // 当前的升级进度
	Module      string    `gorm:"column:module"`        //模块名称
	SrcVersion  string    `gorm:"column:src_version"`   // 设备的原固件版本
	TaskDesc    string    `gorm:"column:task_desc" `    // 升级作业描述信息
	TaskStatus  int       `gorm:"column:task_status"`   // 设备升级状态
	Timeout     string    `gorm:"column:timeout"`       // 设备升级超时时间，单位是分钟
	UtcCreate   time.Time `gorm:"column:utc_create"`    // 升级作业创建时的时间，UTC格式
	UtcModified time.Time `gorm:"column:utc_modified"`  // 升级作业最后一次修改时的时间，UTC格式
	stores.Time
}

func (m *DmOtaUpgradeTask) TableName() string {
	return "dm_ota_upgrade_task"
}

type DmOtaModule struct {
	ID            int64  `gorm:"column:id;type:BIGINT;primary_key;AUTO_INCREMENT"`
	ModuleName    string `gorm:"column:module_name"`    // OTA模块名称，产品下唯一且不可修改。仅支持英文字母、数字、英文句号（.）、短划线（-）和下划线（_）。长度限制为1~64个字符。
	ProductId     string `gorm:"column:product_id"`     // OTA模块所属产品的ProductId。
	ModuleVersion string `gorm:"column:module_version"` //模块版本
	AliasName     string `gorm:"column:alias_name"`     // OTA模块别名。支持中文、英文字母、数字、英文句号（.）、短划线（-）和下划线（_），长度限制为1~64个字符。
	Desc          string `gorm:"column:desc"`           // OTA模块的描述信息，支持最多100个字符。
	stores.Time
}

func (m *DmOtaModule) TableName() string {
	return "dm_ota_module"
}