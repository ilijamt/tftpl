Terraform Templater
===================

Installation
------------

```bash
go install github.com/ilijamt/tftpl/cmd/tftpl
```

Functions
---------

* **upper** - Convert the text to upper case
```gotemplate
{{ "i am lowercase" | upper }}
```

* **lower** - Convert the text to lower cause
```gotemplate
{{ "I AM UPERCASE" | lower }}
```
* **replace** - Replace text 

So if **Value = "demo.replace-dash"**

```gotemplate
{{ replace .Value "_" "-" "." }}
```

We will get **"demo_replace_dash"**

* **default** - Returns a default value if input is empty

```gotemplate
{{ default .Value .Default }}
```