// Custom AIT Portfolio JavaScript
jQuery(document).ready(function ($) {
    console.log('AIT Portfolio Custom JS Loaded');

    // Initialize Isotope for portfolio grid filtering if needed
    if ($('.uc_post_grid_style_one_wrap').length) {
        var $grid = $('.uc_post_grid_style_one_wrap').isotope({
            itemSelector: '.uc_post_grid_style_one_item',
            layoutMode: 'fitRows'
        });

        // Layout Isotope after each image loads
        $grid.imagesLoaded().progress(function () {
            $grid.isotope('layout');
        });
    }

    // Smooth scrolling for anchor links
    $('a[href^="#"]').on('click', function (e) {
        var target = $(this.getAttribute('href'));
        if (target.length) {
            e.preventDefault();
            $('html, body').stop().animate({
                scrollTop: target.offset().top - 100
            }, 1000);
        }
    });

    // Add active class to navigation on scroll
    $(window).on('scroll', function () {
        var scrollPos = $(window).scrollTop() + 150;

        $('section').each(function () {
            var currLink = $(this);
            var refElement = currLink;

            if (refElement.position() && refElement.position().top <= scrollPos && refElement.position().top + refElement.height() > scrollPos) {
                $('a[href="#' + currLink.attr('id') + '"]').addClass('active').parent().siblings().find('a').removeClass('active');
            }
        });
    });
});
