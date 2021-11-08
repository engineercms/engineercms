module github.com/engineercms/engineercms

go 1.14

require (
	github.com/3xxx/flow v0.9.0
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/astaxie/beego v1.12.3
	github.com/beego/admin v0.0.0-20210305083807-6b74f2e7468f
	github.com/bitly/go-simplejson v0.5.0
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/boombuler/barcode v1.0.1
	github.com/casbin/beego-orm-adapter v1.0.0
	github.com/casbin/casbin v1.7.0
	github.com/casbin/xorm-adapter v1.0.0
	github.com/dchest/lru v0.0.0-20151022103600-d8fd1e40a385 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20211104170603-75263a5e99d2
	github.com/eyebluecn/tank v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-xorm/xorm v0.7.9 // indirect
	github.com/golang-jwt/jwt v3.2.1+incompatible
	github.com/google/go-tika v0.2.0
	github.com/gorilla/websocket v1.4.2
	github.com/holys/initials-avatar v0.0.0-20180809162153-a82edcad3408
	github.com/howeyc/fsnotify v0.9.0
	github.com/jinzhu/gorm v1.9.16
	github.com/kardianos/service v1.2.0
	github.com/lib/pq v1.10.1
	github.com/lifei6671/gocaptcha v0.2.0
	github.com/lifei6671/mindoc v1.0.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/pborman/uuid v1.2.1
	github.com/pdfcpu/pdfcpu v0.3.11
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/smartystreets/goconvey v1.6.6
	github.com/tealeg/xlsx v1.0.5
	github.com/unidoc/unioffice v1.10.0
	golang.org/x/crypto v0.0.0-20201002170205-7f63de1d35b0
	golang.org/x/text v0.3.6
	gopkg.in/asn1-ber.v1 v1.0.0-20181015200546-f715ec2f112d // indirect
	gopkg.in/ldap.v2 v2.5.1
	stathat.com/c/consistent v1.0.0 // indirect
	xorm.io/xorm v1.0.7
)

replace github.com/eyebluecn/tank => ./pkg/tank

replace github.com/pdfcpu/pdfcpu => ./pkg/pdfcpu
