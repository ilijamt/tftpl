Examples
========

Google DNS Record Set
---------------------

* [Template file](google_dns_record_set/tpl)
* [Data](google_dns_record_set/data.yml)

Running the **tftpl** command will generate this code that can be run by terraform

```bash
$ tftpl build google_dns_record_set/tpl -f google_dns_record_set/data.yml

resource "google_dns_record_set" "example_com_localhost" {
  provider     = "google.primary-dns"
  managed_zone = "example-com"
  ttl          = "86400"
  name         = "localhost.example.com."
  type         = "A"

  rrdatas = [
    "127.0.0.1",
    "127.0.1.1",
    ]

  depends_on = [
    "google_dns_managed_zone.example_com",
  ]
}

resource "google_dns_record_set" "example_com_www" {
  provider     = "google.primary-dns"
  managed_zone = "example-com"
  ttl          = "14400"
  name         = "www.example.com."
  type         = "CNAME"

  rrdatas = [
    "redirect-me.example.com.",
    ]

  depends_on = [
    "google_dns_managed_zone.example_com",
  ]
}
```