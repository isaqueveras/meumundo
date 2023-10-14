# meumundo
O meumundo é um serviço que tem a funcionalidade de retornar dados referente a uma cidade ou estado do Brasil.

---

## Funcionalidades

Obter todos os dados referente a uma cidade usando apenas uma unica rota.
```json
http://site.com.br/v1/{estado}/{cidade}

{
  "id": "c8546032",
  "city_id": "89e48101",
  "content": "Fortaleza é um município brasileiro, capital do estado do Ceará...",
  "status": "Publish",
  "children": [
    {
      "url": "http://imagem.com/rachel-de-queiroz.png",
      "name": "Rachel de Queiroz",
      "short_desc": "Foi uma tradutora, romancista, escritora, jornalista, cronista prolífica."
    },
    {
      "url": "http://imagem.com/jose-de-alencar.png",
      "name": "José de Alencar",
      "short_desc": "Foi um romancista, dramaturgo, jornalista, advogado e político brasileiro."
    }
  ],
  "border_towns": [
    ["caucaia", "Caucaia", "ce"]
  ]
  "created_at": "2023-09-19T19:47:05.528548-03:00"
}
```

Obter apenas os filhos ilustres de uma cidade
```json
http://site.com.br/v1/{estado}/{cidade}/children

[
  {
    "url": "http://imagem.com/rachel-de-queiroz.png",
    "name": "Rachel de Queiroz",
    "short_desc": "Foi uma tradutora, romancista, escritora, jornalista, cronista prolífica."
  },
  {
    "url": "http://imagem.com/jose-de-alencar.png",
    "name": "José de Alencar",
    "short_desc": "Foi um romancista, dramaturgo, jornalista, advogado e político brasileiro."
  }
]
```

Obter apenas as cidades que fazem fronteira de uma cidade
```json
http://site.com.br/v1/{estado}/{cidade}/border_towns
[
  [
    "aquiraz",
    "Aquiraz",
    "ce"
  ],
  [
    "caucaia",
    "Caucaia",
    "ce"
  ]
]
```

