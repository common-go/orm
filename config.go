package orm

type Config struct {
	MultiStatements bool   `mapstructure:"multi_statements" json:"multiStatements,omitempty" gorm:"column:multistatements" bson:"multiStatements,omitempty" dynamodbav:"multiStatements,omitempty" firestore:"multiStatements,omitempty"`
	DataSourceName  string `mapstructure:"data_source_name" json:"dataSourceName,omitempty" gorm:"column:datasourcename" bson:"dataSourceName,omitempty" dynamodbav:"dataSourceName,omitempty" firestore:"dataSourceName,omitempty"`
	Driver          string `mapstructure:"driver" json:"driver,omitempty" gorm:"column:driver" bson:"driver,omitempty" dynamodbav:"driver,omitempty" firestore:"driver,omitempty"`
	Host            string `mapstructure:"host" json:"host,omitempty" gorm:"column:host" bson:"host,omitempty" dynamodbav:"host,omitempty" firestore:"host,omitempty"`
	Port            int    `mapstructure:"port" json:"port,omitempty" gorm:"column:port" bson:"port,omitempty" dynamodbav:"port,omitempty" firestore:"port,omitempty"`
	Database        string `mapstructure:"database" json:"database,omitempty" gorm:"column:database" bson:"database,omitempty" dynamodbav:"database,omitempty" firestore:"database,omitempty"`
	User            string `mapstructure:"user" json:"user,omitempty" gorm:"column:user" bson:"user,omitempty" dynamodbav:"user,omitempty" firestore:"user,omitempty"`
	Password        string `mapstructure:"password" json:"password,omitempty" gorm:"column:password" bson:"password,omitempty" dynamodbav:"password,omitempty" firestore:"password,omitempty"`
	ConnMaxLifetime int64  `mapstructure:"conn_max_lifetime" json:"connMaxLifetime,omitempty" gorm:"column:connmaxlifetime" bson:"connMaxLifetime,omitempty" dynamodbav:"connMaxLifetime,omitempty" firestore:"connMaxLifetime,omitempty"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns" json:"maxIdleConns,omitempty" gorm:"column:maxidleconns" bson:"maxIdleConns,omitempty" dynamodbav:"maxIdleConns,omitempty" firestore:"maxIdleConns,omitempty"`
	MaxOpenConns    int    `mapstructure:"max_open_conns" json:"maxOpenConns,omitempty" gorm:"column:maxopenconns" bson:"maxOpenConns,omitempty" dynamodbav:"maxOpenConns,omitempty" firestore:"maxOpenConns,omitempty"`
	Mock            bool   `mapstructure:"mock" json:"mock,omitempty" gorm:"column:mock" bson:"mock,omitempty" dynamodbav:"mock,omitempty" firestore:"mock,omitempty"`
	Log             bool   `mapstructure:"log" json:"log,omitempty" gorm:"column:log" bson:"log,omitempty" dynamodbav:"log,omitempty" firestore:"log,omitempty"`
}
