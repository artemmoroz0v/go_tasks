package cache_db
import "gorm.io/driver/postgres"
import "gorm.io/gorm"
import "log"
import "L0/mainstruct"
import "errors"
import "encoding/json"

type Cache struct {
	data map[string]mainstruct.MainStruct
}

type DBdata struct {
	Order_uid  string `gorm:"primaryKey;"`
	Order_json string
}

type DB struct {
	db *gorm.DB
}

func CreateCache() Cache {
	return Cache{data: make(map[string]mainstruct.MainStruct)}
}

func (cache *Cache) Insert(sample mainstruct.MainStruct) error {
	_, ok := cache.data[sample.Order_uid]
	if ok {
		return errors.New("Error: do not repeat infromation in cache!")
	}
	cache.data[sample.Order_uid] = sample
	return nil
}

func (cache *Cache) Get (orderUid string) (mainstruct.MainStruct, error) {
	value, ok := cache.data[orderUid]
	if ok {
		return value, nil
	}
	return mainstruct.MainStruct{}, errors.New("")
}

func (cache *Cache) GetAllCache() []mainstruct.MainStruct {
	slice := make([]mainstruct.MainStruct, len(cache.data), 0)
	for _, order := range cache.data {
		slice = append(slice, order)
	}
	return slice
}

func Connect() (DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=user password=userwb dbname=db_wb sslmode=disable"))
	if err != nil {
		return DB{}, err
	}
	err = db.AutoMigrate(&DBdata{})
	if err != nil {
		return DB{}, err
	}
	return DB{db}, nil
}

func (d *DB) UploadCache(cache *Cache) error {
	var data []DBdata
	err := d.db.Find(&data)
	for _, rec := range data {
		var JsonOrder mainstruct.MainStruct
		err := json.Unmarshal([]byte(rec.Order_json), &JsonOrder)
		if err != nil {
			log.Print(err)
		}
		cache.Insert(JsonOrder)
	}
	return err.Error
}

func (d *DB) InsertInDB(rec...mainstruct.MainStruct) error {
	log.Printf("INSERTING IN DATABASE: %v", rec)
	for _, r := range rec {
		order, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			log.Print(err)
		}
		errGorm := d.db.Create(DBdata{r.Order_uid, string(order)})
		if errGorm.Error != nil {
			return errGorm.Error
		}
	}
	return nil
}
