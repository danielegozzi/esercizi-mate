---
title: "Disegnare le frazioni come quadretti"
date: 2022-10-25T00:00:00+02:00
categories:
- Utility
tags:
- Utility
- Development
---
![Stefania]({{< resource url="dan_stefy.png#floatleft" >}})
Questa è la terza stesura del programma che disegna la frazioni come griglie di quadretti colorati.

Dalla prima versione in Ruby, generando l'immagine sul server, sono passato a riscrivere in linguaggio Go lo stesso programma. Ora è tutto fatto in Javascript nel browser stesso.

Dedicato a Stefania, mia moglie, che lo usa per scrivere le sue verifiche.
<!--more-->

{{< unsafe >}}
<script defer src="/esercizi-mate/frac.js"></script>
<hr>
<div style="clear: both">
<canvas id="canvas" width="200" height="200"></canvas>
<br><a href="" id="download_link">Download</a>
</div>
{{< /unsafe >}}

<form id="fraction_squares" action="" method="GET">
  <label for="scale">Scala di un quadretto (10 px < <em>scala</em> < 30 px)</label> <input type="text" name="scale" size="3" value="25"/> px<br>
  <label for="rows">Numero di righe (max. 20)</label> <input type="text" name="rows" size="3" id="rows" value="7"/><br>
  <label for="columns">Numero di colonne (max. 20)</label> <input type="text" name="columns" size="3" id="cols" value="3"/><br>
  <label for="filled">Quadretti pieni</label> <input type="text" name="filled" size="2" value="5"/>
  <input type="submit" value="Disegna la frazione"/>
</form>

{{< unsafe >}}
</div>
{{< /unsafe >}}
