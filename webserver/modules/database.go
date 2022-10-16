package modules

import (
	"context"
	"log"
	"site/webserver/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	conn *pgxpool.Pool
	tx  pgx.Tx
}

// создает новое соединение postgres
func NewPostgres(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return conn, nil
}


// NewDatabaseInstance - конструктор структуры экземпляра базы данных
func NewDatabase(conn *pgxpool.Pool) *DB {
	return &DB{
		conn: conn,
	}
}

func (db *DB) Get(ctx context.Context) ([]models.User,error) {
	return []models.User{},nil
}
func (db *DB) SetUser(ctx context.Context, user models.User) (uint64,error){
	sql := `INSERT INTO users (firstname, lastname, displayname, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	row := db.conn.QueryRow(context.Background(),sql,
		user.FirstName,
		user.LastName,
		user.DisplayName,
		user.Email,
		user.Password,
	)
	var id uint64
 	err := row.Scan(&id)
 	if err != nil {
 		log.Println(err)
 		return 0, err
 	}
 	return id, nil
}

// // получение всех order
// func (db *DB) Get(ctx context.Context) ([]models.Order,error) {

// 	sql := `SELECT * FROM "order"`
// 	rows, err := db.conn.Query(ctx,sql)
// 	if err != nil {
// 		return nil,err
// 	}
// 	var data []models.Order

// 	for rows.Next() {

// 		var row models.Order

// 		values, err := rows.Values()
// 		if err != nil {
// 			return nil, err
// 		}
// 		// todo
// 		row.OrderUID = values[0].(string)
// 		row.TrackNumber = values[1].(string)
// 		row.Entry = values[2].(string)
// 		row.Delivery = models.Delivery{}
// 		if err := json.Unmarshal([]byte(values[3].(string)), &row.Delivery); err != nil {
//         	return nil,err
//     	}
// 		row.Payment = models.Payment{}
// 		if err := json.Unmarshal([]byte(values[4].(string)), &row.Payment); err != nil {
//         	return nil,err
//     	}
// 		row.Items = []models.Item{}
// 		if err := json.Unmarshal([]byte(values[5].(string)), &row.Items); err != nil {
//         	return nil,err
//     	}
// 		row.Locale = values[6].(string)
// 		row.InternalSignature = values[7].(string)
// 		row.CustomerID = values[8].(string)
// 		row.DeliveryService = values[9].(string)
// 		row.Shardkey = values[10].(string)
// 		row.SmID = values[11].(int32)
// 		row.DateCreated = values[12].(string)
// 		row.OofShard = values[13].(string)

// 		data = append(data,row)

// 	}

// 	return data,nil

// }

// Добавление order
// func (db *DB) Set(ctx context.Context, order models.Order) (string,error){

	// delivery, err := json.Marshal(order.Delivery)
	// if err != nil {
	// 	return "", err
	// }
	// payment, err := json.Marshal(order.Payment)
	// if err != nil {
	// 	return "", err
	// }
	// items, err := json.Marshal(order.Items)
	// if err != nil {
	// 	return "", err
	// }

	// sql := `INSERT INTO "order" (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
 //        VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING order_uid;`
 //    row := db.conn.QueryRow(ctx,sql,
 //    	order.OrderUID,
 //    	order.TrackNumber,
 //    	order.Entry,
 //    	delivery,
 //    	payment,
 //    	items,
 //    	order.Locale,
 //    	order.InternalSignature,
 //    	order.CustomerID,
 //    	order.DeliveryService,
 //    	order.Shardkey,
 //    	order.SmID,
 //    	order.DateCreated,
 //    	order.OofShard,
 //    )
 //    var order_uid string
 //  	err = row.Scan(&order_uid)
 //  	if err != nil {
 //  		return "",err
	// }
	// return order_uid,nil

// }