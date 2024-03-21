package main

import (
	"github.com/ziterz/shopeefresh-server-go/initializers"
	"github.com/ziterz/shopeefresh-server-go/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

var products = []models.Product{
	{
		Name:  "Buah Beet Root 1 kg",
		Image: "https://down-id.img.susercontent.com/file/sg-11134201-7rbku-ln4i7t9u6rnv4a_tn",
		Stock: 100,
		Price: 17500,
	},
	{
		Name:  "Cabai Rawit Merah 500 gr",
		Image: "https://down-id.img.susercontent.com/file/ba00d1c1d34684b8f4e60556182a294b_tn",
		Stock: 100,
		Price: 17000,
	},
	{
		Name:  "Bawang Merah Brebes 1 kg",
		Image: "https://down-id.img.susercontent.com/file/id-11134207-7qul5-ljnmo5zkc7yb15_tn",
		Stock: 100,
		Price: 22500,
	},
	{
		Name:  "Buncis 500 gr",
		Image: "https://down-id.img.susercontent.com/file/ebf2395fb4611b91236d785f059ef4fc_tn",
		Stock: 100,
		Price: 0,
	},
	{
		Name:  "Daun Kemangi 100 gr",
		Image: "https://down-id.img.susercontent.com/file/53f6da224fae753c0a714d7d08e5c611_tn",
		Stock: 100,
		Price: 5000,
	},
	{
		Name:  "Cabai Hijau Keriting 500 gr",
		Image: "https://down-id.img.susercontent.com/file/ff59a4c9f27d50c77ff29975bc5c8946_tn",
		Stock: 100,
		Price: 17000,
	},
	{
		Name:  "Wortel Berastagi 500 gr",
		Image: "https://down-id.img.susercontent.com/file/31800cfb6d6247437868612653c8b9f7_tn",
		Stock: 100,
		Price: 11000,
	},
	{
		Name:  "Timun Segar 500 gr",
		Image: "https://down-id.img.susercontent.com/file/bb6661466dc8da797acefca7e6cedd04_tn",
		Stock: 100,
		Price: 5000,
	},
	{
		Name:  "Sawi Hijau Caisim 500 gr",
		Image: "https://down-id.img.susercontent.com/file/cae2bf2e6f3a46376ff1e4a7bc4bfbb7_tn",
		Stock: 100,
		Price: 8000,
	},
	{
		Name:  "Daun Jeruk per Pack",
		Image: "https://down-id.img.susercontent.com/file/a731c3fb5f3e84ed3cd921e515ab1649_tn",
		Stock: 100,
		Price: 5000,
	},
	{
		Name:  "Bawang Bombai Import 500 gr",
		Image: "https://down-id.img.susercontent.com/file/id-11134207-7qul2-lkc2l6f5c4gp1b_tn",
		Stock: 100,
		Price: 12500,
	},
	{
		Name:  "Kentang Dieng Fresh 1 kg",
		Image: "https://down-id.img.susercontent.com/file/id-11134207-7qukz-lkht6o9vyr0g52_tn",
		Stock: 100,
		Price: 14500,
	},
	{
		Name:  "Terong Ungu 500 gr",
		Image: "https://down-id.img.susercontent.com/file/de43158b067b2c457b89b647dbe475c2_tn",
		Stock: 100,
		Price: 7000,
	},
	{
		Name:  "Baby Pakcoy 250 gr",
		Image: "https://down-id.img.susercontent.com/file/ac56533929cef994d46eee6225217bc8_tn",
		Stock: 100,
		Price: 4800,
	},
	{
		Name:  "Tomat Hijau 500 gr",
		Image: "https://down-id.img.susercontent.com/file/a4312d6992404cfbb7e16f30d7e528ff_tn",
		Stock: 100,
		Price: 9000,
	},
}

func main() {
	initializers.DB.Create(&products)
}
