# Movie Festival App

Menggunakan golang dan framework gin dan gorm

## API dapat menyimpan Movie

API yang Anda buat harus dapat menyimpan movie melalui route :

- Method : POST

- URL : /movies

- Body Request:

```
{
    "Title": string,
    "Description": string,
    "Duration" : int,
    "Artist" : string,
    "Genre" : string,
    "URL" : string,
}
```

Response to be returned :

- Status Code : 200
- Response Body :

``` json
{
    "data": {
        "id": 17,
        "UUID": "a84befef-7f4c-49ec-97a7-b2c59759f1e9",
        "title": "Testing",
        "description": "Testing",
        "duration": 200,
        "artist": "John Lennon",
        "genre": "Comedy",
        "url": "http://localhost:8080/movies",
        "viewers": 0,
        "created_at": "2023-02-10T14:43:49.727621+07:00",
        "updated_at": "2023-02-10T14:43:49.727621+07:00"
    }
}
```

## API dapat menampilkan seluruh movies

API yang Anda buat harus dapat menampilkan semua Movie yang disimpan melalui route :

- Method : GET

- URL : /movies/list

## API dapat menampilkan satu movie

API yang Anda buat harus dapat menampilkan Movie yang disimpan melalui route :

- Method : GET

- URL : /movies/:uuid

## API dapat mengubah data movie

API yang Anda buat harus dapat mengubah data Movie berdasarkan id melalui rute:

- Method : PATCH

- URL : /movies/:uuid

- Body Request:

```
{
    "Title": string,
    "Description": string,
    "Duration" : int,
    "Artist" : string,
    "Genre" : string,
    "URL" : string,
}
```
