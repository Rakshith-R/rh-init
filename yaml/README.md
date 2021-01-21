# YAML : YAML Ain't Markup Language
Popularly used for writing configs.!

https://rollout.io/blog/yaml-tutorial-everything-you-need-get-started/

https://codebeautify.org/yaml-to-json-xml-csv
```yaml
---
#begins with --- , end with ... , space used for scoping
key : value
text1 : hello
text2 : "hello"
text3 : > #folds back replaces newline with space
 hello how are you
 hello
text4 : |  #add '+' to preserve last characters, '-' to strip ws or nlines
 hello how are you
 hello
list1 : 
 - 123
 - 0123  #oct
 - 0xA  #hex
 - null
 - ~
list2 : &Anchor [1,"1",4]   #Node
list3 : *Anchor #unpacks here
nested_object :
 inside : {embedded_json:TRUE}
 inside2 : "tEST"
...
```