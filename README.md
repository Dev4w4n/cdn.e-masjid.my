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
    "base64_file" : "iVBORw0KGgoAAAANSUhEUgAAAMAA...."
}
```
File to Base64 reference: https://base64.guru/converter/encode/file

### Allowed Mime types
```
image/gif
image/jpeg
image/png
image/webp
application/pdf
```

### Response

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
    "path": "https://cdn.e-masjid.my/demo/951ef9cf-eab5-40d1-ad24-5afa2d2de185.jpg",
    "created_at": "65586126"
}
```
</td>
</tr>
<tr>
<td> 403 </td>
<td>

**Forbidden**

</td>
</tr>

<tr>
<td> 500 </td>
<td>

**Internal server error**

</td>
</tr>

</table>