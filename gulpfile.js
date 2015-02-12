var gulp = require('gulp');
var watch = require('gulp-watch');
var gutil = require('gulp-util');
var source = require('vinyl-source-stream');
var buffer = require('vinyl-buffer');
var plumber  = require('gulp-plumber');
var streamify = require('gulp-streamify');

var watchify = require('watchify');
var browserify = require('browserify');
var uglify = require('gulp-uglify');
var sourcemaps = require('gulp-sourcemaps');

var jade = require('gulp-jade');

var _ = require('lodash');

gulp.task('js', function() {
  var FILES = ["main"];
  watchify.args.fullPaths = false;
  _.each(FILES, function(file) {
    var bundler = watchify(browserify('./source/js/' + file + '.coffee', { cache: {}, packageCache: {}, fullPaths: true, debug: true }));

    bundler.on('update', rebundle);

    function rebundle() {
      return bundler.bundle()
        .on('error', gutil.log.bind(gutil, 'Browserify Error'))
        .pipe(source(file + '.js'))
        .pipe(buffer())
        .pipe(sourcemaps.init({loadMaps: true}))
        .pipe(streamify(uglify()))
        .pipe(sourcemaps.write('./'))
        .pipe(gulp.dest('./public/js/'));
    }

    return rebundle();
  });

});

gulp.task('jade', function () {
  gulp.src('source/jade/**/*.jade')
    .pipe(watch('source/jade/**/*.jade'))
    .pipe(jade({pretty: true}))
    .pipe(gulp.dest('./app/views/'))
});

gulp.task('watch', ['js', 'jade']);
