(function() {
    function getUrlParameterByName(name, url) {
         if (!url) url = window.location.href;
         name = name.replace(/[\[\]]/g, "\\$&");
         var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"), results = regex.exec(url);
         if (!results) return null;
         if (!results[2]) return '';
         return decodeURIComponent(results[2].replace(/\+/g, " "));
    }
    function makeChart(svg) {
        var chart = new dimple.chart(svg);
        var x = chart.addCategoryAxis("x", "date");
        x.title = "Date";
        x.timeField = "date";
        var y = chart.addMeasureAxis("y", "stars");
        y.title = "Popularity"
        chart.addLegend(60, 10, 500, 20, "right");
        window.onresize = function() {
            chart.draw(0, true);
        };
        return chart;
    }
    function populateChart(chart, repositories) {
        repositories.forEach(function(repository) {
            d3.json("/api/stars/" + repository, function (starsData) {
                addChartSeries(chart, starsData, repository);
            });
        })
    }
    function addChartSeries(chart, data, title) {
        var series = chart.addSeries(title, dimple.plot.line);
        series.interpolation = "cardinal";
        series.data = data;
        chart.draw(1000);
    }

    var repositories = getUrlParameterByName("repositories").split(",")
    var svg = dimple.newSvg("#chartContainer", "90%", "90%");
    var chart = makeChart(svg);

    populateChart(chart, repositories)
})()