var starsApi = "/api/stars"
var repository = "mlucchini/githubcompare"
var url = starsApi + "/" + repository

var svg = dimple.newSvg("#chartContainer", 590, 400);
d3.json(url, function (data) {
  var myChart = new dimple.chart(svg, data);
  myChart.setBounds(60, 30, 505, 305);
  myChart.addCategoryAxis("x", "date");

  myChart.addMeasureAxis("y", "stars");

  var s = myChart.addSeries(repository, dimple.plot.line);
  s.interpolation = "cardinal";

  myChart.addLegend(60, 10, 500, 20, "right");
  myChart.draw();
});