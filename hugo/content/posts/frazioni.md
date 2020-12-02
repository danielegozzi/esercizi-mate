---
title: "Disegnare le frazioni come quadretti"
date: 2020-12-02T00:00:00+01:00
categories:
- Utility
tags:
- Utility
- Development
---
Dedico a mia moglie il mio primo tentativo di rifare con il linguaggio go una applicazione semplice che disegna le frazioni per le sue verifiche.

<!--more-->

<form id="fraction_squares" action="/frac" method="GET">
  <label for="scale">Scala di un quadretto (10 px < <em>scala</em> < 30 px)</label> <input type="text" name="scale" size="3"/> px<br>
  <label for="rows">Numero di righe (max. 20)</label> <input type="text" name="rows" size="3" id="rows"/><br>
  <label for="columns">Numero di colonne (max. 20)</label> <input type="text" name="columns" size="3" id="cols"/><br>
  <label for="filled">Quadretti pieni</label> <input type="text" name="filled" size="2"/>
  <input type="submit" value="Disegna la frazione"/>
</form>

