Examples
========

Google DNS Record Set
---------------------

* [Template file](google_dns_record_set/tpl)
* [Data](google_dns_record_set/data.yml)

Running the **tftpl** command will generate this code that can be run by terraform

```bash
$ tftpl build google_dns_record_set/tpl -f google_dns_record_set/data.yml
```

```text
resource "google_dns_record_set" "example_com_type_a_main" {
  provider     = "google.primary-dns"
  managed_zone = "example-com"
  ttl          = "3600"
  type         = "A"
  name         = "example.com."

  rrdatas = [
    "1.1.1.1",
    "1.1.1.2",
  ]

  depends_on = [
    "google_dns_managed_zone.example_com",
  ]
}

resource "google_dns_record_set" "example_com_type_a_localhost" {
  provider     = "google.primary-dns"
  managed_zone = "example-com"
  ttl          = "86400"
  type         = "A"
  name         = "localhost.example.com."

  rrdatas = [
    "127.0.0.1",
    "127.0.1.1",
  ]

  depends_on = [
    "google_dns_managed_zone.example_com",
  ]
}

resource "google_dns_record_set" "example_com_type_cname_www" {
  provider     = "google.primary-dns"
  managed_zone = "example-com"
  ttl          = "3600"
  type         = "CNAME"
  name         = "www.example.com."

  rrdatas = [
    "redirect-me.example.com.",
  ]

  depends_on = [
    "google_dns_managed_zone.example_com",
  ]
}
```