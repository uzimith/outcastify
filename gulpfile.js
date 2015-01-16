var gulp = require('gulp');
var gutil = require('gulp-util');
var source = require('vinyl-source-stream');
var plumber  = require('gulp-plumber');
var watchify = require('watchify');
var browserify = require('browserify');
var sass     = require('gulp-ruby-sass');
var pleeease = require('gulp-pleeease');
var _ = require('lodash');

gulp.task('js-watch', function() {
  // watchify
  var FILES = ["main"];
  watchify.args.fullPaths = false;
  _.each(FILES, function(file) {
    var bundler = watchify(browserify('./source/js/' + file + '.js', watchify.args));

    bundler.transform('6to5ify')

    bundler.on('update', rebundle);

    function rebundle() {
      return bundler.bundle()
        .on('error', gutil.log.bind(gutil, 'Browserify Error'))
        .pipe(source(file + '.js'))
        .pipe(gulp.dest('./public/js/'));
    }

    return rebundle();
  });

});

gulp.task('sass', function() {
  gulp.src('source/sass/**/*.sass')
    .pipe(plumber())
    .pipe(sass({
        style: 'nested',
        compass: true
    }))
    .pipe(pleeease({
        autoprefixer: {
            browsers: ['last 2 versions']
        },
        minifier: false
    }))
    .pipe(gulp.dest('public/css/'));
});

gulp.task('sass-watch', function () {
  gulp.watch('source/sass/**/*.sass', ['sass']);
})
