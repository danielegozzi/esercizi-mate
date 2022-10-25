var canvas;
var ctx = null;
var filename = "frac.png"

window.addEventListener("DOMContentLoaded", function() {
  canvas = document.getElementById('canvas');
  if (canvas.getContext) {
    ctx = canvas.getContext('2d');
  }
  document.getElementById('fraction_squares').addEventListener("submit", function(e) {
    e.preventDefault(); // before the code
    drawGrid();
  });
  document.getElementById('download_link').addEventListener("click", function(e) {
    e.preventDefault(); // before the code
    downloadCanvas();
  });
  drawGrid();
});

function downloadCanvas() {
  var link = document.createElement('a');
  link.setAttribute('download', filename);
  link.setAttribute('href', canvas.toDataURL("image/png").replace("image/png", "image/octet-stream"));
  link.click();
}

function drawGrid() {
  canvas = document.getElementById('canvas');
  ctx = canvas.getContext('2d');
  scale = parseInt(document.querySelector("input[name=scale]").value);
  rows = parseInt(document.querySelector("input[name=rows]").value);
  columns = parseInt(document.querySelector("input[name=columns]").value);
  filled = parseInt(document.querySelector("input[name=filled]").value);
  canvas.width = (scale-ctx.lineWidth)*columns + ctx.lineWidth*2;
  canvas.height = (scale-ctx.lineWidth)*rows + ctx.lineWidth*2;
  filename = "frac_scale" + scale.toString() + "_rows" + rows.toString() + "_cols" + columns.toString() + "_filled" + filled.toString() + ".png"
  ctx.lineWidth = 1;
  if (scale > 15) {
    ctx.lineWidth = 2;
  }
  ctx.fillStyle = "rgb(128, 128, 128)";
  var fillCounter = 1;
  for (let x=0; x<columns; x++) {
    for (let y=0; y<rows; y++) {
      if (fillCounter <= filled) {
        ctx.fillRect(x*(scale-ctx.lineWidth)+ctx.lineWidth, y*(scale-ctx.lineWidth)+ctx.lineWidth, scale-ctx.lineWidth, scale-ctx.lineWidth);
        fillCounter++;
      }
      ctx.strokeRect(x*(scale-ctx.lineWidth)+ctx.lineWidth, y*(scale-ctx.lineWidth)+ctx.lineWidth, scale-ctx.lineWidth, scale-ctx.lineWidth);
    }
  }
}
