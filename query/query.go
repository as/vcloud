package query

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"time"

	"github.com/as/vcloud"
	"github.com/as/vcloud/util"
)

const (
	//queryUriFmt string = "https://%s/api/query/?type=%s&pageSize=%d&sortDesc=%s" // Request URL for a Query
	queryUriFmt string = "https://%s/api/query/?type=%s" // Request URL for a Query
)

type Links []Link

type Date string

func (d Date) String() (r string) {
	t, err := time.Parse(time.RFC3339, string(d))
	t = t.Local()
	if err != nil {
		r = fmt.Sprintf("ERROR_PARSING")
	}
	r = fmt.Sprintf("%04v-%02d-%02v_%02v:%02v:%02v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return r
}

func (l Links) HrefOf(rel string) string {
	for _, v := range l {
		if v.Rel == rel {
			return v.Href
		}
	}
	return ""
}




func NewOptions() (o *Options) {
	o = new(Options)
	o.PageSize = 50
	o.Limit = 1000
	return o
}

type Options struct {
	Page     int
	NoPages  int
	PageSize int
	Limit    int
	Filter   string
	Href     string
	Sort     string
	Element  interface{}
}

func (q *Options) Validate() {
	if q.PageSize > q.Limit {
		q.PageSize = q.Limit
	}

	if q.PageSize == 0 {
		q.PageSize = 100
	}

	if q.Sort == "" {
		q.Sort = "creationDate"
	} else {
		q.Sort = util.C9toAPI(q.Sort)
	}

}

func (q *Options) makeUrl(s *vcloud.Session) (string, error) {
	estr := reflect.TypeOf(q.Element).Name()

	url := fmt.Sprintf(queryUriFmt, s.Server, UriParams[estr])

	if q.PageSize != 0 {
		url += fmt.Sprintf("&%s=%d", "pageSize", q.PageSize)
	}

	if q.Sort != "" {
		url += fmt.Sprintf("&%s=%s", "sortDesc", q.Sort)
	}

	if q.Filter != "" {
		url += fmt.Sprintf("&%s=%s", "filter", q.Filter)
	}

	return url, nil
}

func (q *Options) Url(s *vcloud.Session) (string, error) {
	var url string

	q.Validate()

	switch e := q.Element.(type) {
	case string:
		url = e
	case ApiDefinitionRecord, CatalogItemRecord, CatalogRecord, DiskRecord, EventRecord,
		FileDescriptorRecord, GroupRecord, MediaRecord, OrgNetworkRecord, OrgVdcRecord,
		OrgVdcStorageProfileRecord, ResultRecords, ServiceRecord, TaskRecord,
		UserRecord, VAppNetworkRecord, VAppRecord, VAppTemplateRecord, VMRecord,
		VmDiskRelationRecord:

		//estr := reflect.TypeOf(e).Name()
		//queryUriFmt string = "https://%s:%s/api/query/?type=%s&pageSize=%d&sortDesc=%s" // Request URL for a Query
		url, _ = q.makeUrl(s)
		//url = fmt.Sprintf(queryUriFmt, s.Host, s.Port, UriParams[estr], q.PageSize, q.Sort)
	case Link:
		url = e.Href
	default:
		return "", fmt.Errorf("Error %v\n", e)
	}

	return url, nil
}

// Method Find is similar to Query, except it concatenates all of the
// Records in multiple ResultRecords structures
func FullQuery(s *vcloud.Session, o *Options) (interface{}, error) {
	opts := *o
	//TODO: Process options

	qr, err := Query(s, opts)
	if err != nil {
		return nil, err
	}

	if qr.Total == 0 {
		return nil, nil
	}

	//TODO: Error checking here, qr should be a slice
	dst := reflect.ValueOf(qr.Records)
	ps := opts.PageSize

	var i int // Number of records (pages * pageSize)

	for i = ps; opts.Limit == 0 || i < opts.Limit; i += ps {
		opts.Href = qr.Links.HrefOf("nextPage")

		if opts.Href == "" {
			break
		}

		//TODO: Error checking here, qr should be a slice
		qr, err = Query(s, opts)
		if err != nil {
			panic(err) //TODO
		}

		dst = reflect.AppendSlice(dst, reflect.ValueOf(qr.Records))
	}

	if opts.Limit != 0 && i > opts.Limit && opts.Limit < qr.Total {
		dst = dst.Slice(0, opts.Limit)
	}

	return dst.Interface(), nil
}

// Query executes a vCloud query based on element's type
// and returns a Record interface. Note: The reason "Record" isn't
// embedded into ResultRecords is due to a bug in Go 1.2.1 where
// it is impossible to Unmarshal() into an interface
func Query(s *vcloud.Session, opts Options) (*ResultRecords, error) {
	var (
		url string
		qr  ResultRecords
	)

	// opts.Href overrides a query URL
	if opts.Href != "" {
		url = opts.Href
	} else {
		url, _ = opts.Url(s)
	}

	body, err := s.DoRequestGetBody("GET", url, nil)
	fmt.Println(string(body))
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &qr)
	if err != nil {
		return nil, err
	}

	// Assigns the Record interface to the actual Record struct
	// obtained from running the Query. Switch is
	// ordered by most-commonly-requested Query record
	switch opts.Element.(type) {
	case VAppTemplateRecord:
		qr.Records = qr.VAppTemplateRecords
	case VAppRecord:
		qr.Records = qr.VAppRecords
	case VAppNetworkRecord:
		qr.Records = qr.VAppNetworkRecords
	case OrgVdcRecord:
		qr.Records = qr.OrgVdcRecords
	case UserRecord:
		qr.Records = qr.UserRecords
	case CatalogItemRecord:
		qr.Records = qr.CatalogItemRecords
	case CatalogRecord:
		qr.Records = qr.CatalogRecords
	case VMRecord:
		qr.Records = qr.VMRecords
	case EventRecord:
		qr.Records = qr.EventRecords
	case TaskRecord:
		qr.Records = qr.TaskRecords
	case OrgNetworkRecord:
		qr.Records = qr.OrgNetworkRecords
	case ApiDefinitionRecord:
		qr.Records = qr.ApiDefinitionRecords
	case DiskRecord:
		qr.Records = qr.DiskRecords
	case FileDescriptorRecord:
		qr.Records = qr.FileDescriptorRecords
	case GroupRecord:
		qr.Records = qr.GroupRecords
	case MediaRecord:
		qr.Records = qr.MediaRecords
	case OrgVdcStorageProfileRecord:
		qr.Records = qr.OrgVdcStorageProfileRecords
	case ServiceRecord:
		qr.Records = qr.ServiceRecords
	case VmDiskRelationRecord:
		qr.Records = qr.VmDiskRelationRecords
	}
	fmt.Printf("%#v\n",qr)
	return &qr, nil
}

