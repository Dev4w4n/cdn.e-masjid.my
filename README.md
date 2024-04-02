# E-Masjid.My CDN-API  Documentation

[ Base URL: https://cdn.e-masjid.my ]

Allowed origins: https://*.e-masjid.my

## POST /api/upload
Upload a file to cdn.e-masjid.my

### Parameter
``` json

body (object)
{
    "mime_type" : "image/jpeg",
    "sub_domain" : "localhost",
    "table_reference" : "tetapan",
    "mark_as_delete" : false,
    "data" : "[in bytes]"
}
```

### Allowed Mime types
```
image/gif
image/jpeg
image/png
image/webp
application/pdf
```

## Response

<table>
<tr>
<td> Status </td> <td> Response </td>
</tr>
<tr>
<td> 201 </td>
<td>

```json
application/json
{
    "id": 10,
    "path": "https://cdn.e-masjid.my/files/213123-e21ed-12e-12e--12e.jpg",
    "created_at": "2024-04-01T20:45:26.433Z"
}
```
</td>
</tr>
<tr>
<td> 401 </td>
<td>

**Unauthorized**

</td>
</tr>

<tr>
<td> 500 </td>
<td>

**Internal server error**

</td>
</tr>

</table>