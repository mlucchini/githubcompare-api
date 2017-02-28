# GitHub Compare API

GitHub Compare API lets you compare the popularity of multiple repositories over time.
This service is fed by data generated with [GitHub Compare](https://github.com/mlucchini/githubcompare).

The service is deployed on [Google App Engine](http://githubcompare.appspot.com/#repositories=angular/angular.js,facebook/react,jashkenas/backbone,emberjs/ember.js)
but has not been kept up-to-date. It uses Google Datastore, Tasks and Google Storage.

## Main usage

```sh
make
make serve
```

## GitHub Compare UI

The UI is made with React and React Toolbox for material design UI components.
Comparison of container technologies over time:

<p align="center">
  <img alt="Comparison of container technologies over time" src="https://github.com/mlucchini/githubcompare-api/blob/master/README.png" width="80%" />
</p>
