package adapter

import (
        "log"
        "database/sql"
)
import (
        _ "github.com/go-sql-driver/mysql"
        "encoding/xml"
)

var db *sql.DB;

type Keys struct {
        Guid []string `xml:"guid"`
}

type LastSessionTime struct {
        LastTime string `xml:"lastSessionDateTime"`
}

type Invoice struct {
        XMLName xml.Name `xml:"GetChangesResult"`
        id int
        status int
        ordersCount int
        sum float32
        Uuid string `xml:"Guid"`
}

func GetChanges (LastTime interface{}) []byte {
        var invoices []*Invoice
        rows, _ := db.Query(getChangesQuery(LastTime.(*LastSessionTime).LastTime))
        for rows.Next() {
                i := new(Invoice)
                rows.Scan(&i.id, &i.status, &i.ordersCount, &i.sum, &i.Uuid)
                invoices = append(invoices, i)
        }
        defer rows.Close()

        return getXML(invoices)
}

func ReadIntegrationMessages(guid interface{}) []byte {
        log.Println(guid)
        for _, gd := range guid.(*Keys).Guid {
                log.Println(gd)
        }
        return []byte("")
}

func getChangesQuery(lastTime string) string {
        sql := `SELECT j.id, j.status, count(DISTINCT bo2.id) as 'orders', bp.amount, UUID() as guid FROM documents AS j
        LEFT JOIN doc2package AS cdp ON (j.id = cdp.document_id)
        LEFT JOIN packages AS cp ON (cp.id = cdp.package_id)
        LEFT JOIN orders AS bo ON (bo.group_id = cp.orders_group_id)
        LEFT JOIN orders AS bo2 ON (cp.orders_group_id = bo2.group_id)
        LEFT JOIN payments AS bp ON (bp.group_id = cp.orders_group_id)
        WHERE
        (GREATEST(j.create_date, j.dependent_changed) >= '` + lastTime + `' OR
        bo.mod_date >= '`+lastTime+`') AND (owner_id IN (1254)) AND (`+"`"+`group`+"`"+` = 2)
        AND bo2.state NOT IN (5)`
        return sql
}

func getXML(Entity interface{}) []byte {
	xmlstring, _ := xml.MarshalIndent(Entity, "", "	")
	xmlstring = []byte(xml.Header + string(xmlstring))
	return xmlstring
}

func init ()  {
        db, _ = sql.Open("mysql", "billing:password@tcp(0.0.0.0:3306)/billing")

        //if err != nil {
        //	panic(err.Error())
        //}
        //defer db.Close()
}
