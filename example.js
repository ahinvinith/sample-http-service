import http from 'k6/http';

const CERT = `-----BEGIN CERTIFICATE-----
MIIFgTCCA2kCAQEwDQYJKoZIhvcNAQEFBQAwgYExCzAJBgNVBAYTAlNFMRcwFQYD
VQQIEw5TdG9ja2hvbG1zIExhbjESMBAGA1UEBxMJU3RvY2tob2xtMRcwFQYDVQQK
...
/n5QrTGhP51P9Q1THzRfn6cNCDwzSTMVEJr40QhuTJQWASe3miuFmZoG5ykmGqVm
fWQRiQyM330s9vTwFy14J2Bxe4px6cyy7rVXvYL2LvfA4L0T7/x1nUULw+Mpqun1
R3XRJWqGDjBKXr5q8VatdQO1QLgr
-----END CERTIFICATE-----`;

const KEY = `-----BEGIN RSA PRIVATE KEY-----
KsZVVI1FTX+F959vqu1S02T+R1JM29PkIfJILIXapKQfb0FWrALU5xpipdPYBWp7
j5iSp06/7H8ms87Uz9BrOA6rytoRSE0/wEe5WkWdBBgLLPpfOSWZsAA5RGCB2n+N
...
Dk+frzKuiErHFN7HOHAQannui4eLsY0ehYMByowgJIUGzIJyXR6O19hVhV7Py66u
X7/Jy01JXn83LuWdpaPAKU+B42BLP0IGXt5CocPms07HOdtJ/wm2zwHTyfjn9vu+
HO/dQr6a7DhRu2lLI9Sc983NwRqDKICZQQ/+gqWk8BgQZ1yI9O4AYkzywzAEk3py
-----END RSA PRIVATE KEY-----`;

export const options = {
    tlsAuth: [
        {
            domains: ['example.com'],
            cert: CERT,
            key: KEY,
        },
    ],
};

export default function () {
    http.get('https://example.com/');
}
