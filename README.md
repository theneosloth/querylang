# Querylang


A very simple lexer and parser for a generic search query language.

A query is a series of identifiers that could additionally contain queries in the form of `IDENTIFIER``OPERATOR``QUERY`

For example
```
hp>100 mana=<50 name=Morte 
```

would return the following struct

```
&{[{hp > 100} {mana =< 50} {name = Morte}]}
```
