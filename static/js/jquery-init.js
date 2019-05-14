(function () {
    "use strict";
    $(document).ready(function () {

        // Main mobile navigation
        jQuery(".cs-toggle-main-navigation").on("click", function () {
            jQuery("nav#cs-main-navigation").toggle();
        });

	    // Fitvids
	    jQuery("body").fitVids();

        // Tabs
        jQuery(".cs-tab-group").tabs();

        // Accordions
        jQuery(".cs-accordion-group").accordion({
            heightStyle: "content",
            collapsible: true,
            icons: false
        });

        // Sticky sidebar
        jQuery(".cs-sticky-sidebar").theiaStickySidebar({
            additionalMarginTop: 10
        });

        // Sticky main menu
        jQuery(window).scroll(function () {
            var mainmenu = jQuery(".cs-main-navigation-wrapper.cs-menu-is-sticky");
            if (parseInt(mainmenu.attr("rel"),10) <= Math.abs(parseInt(jQuery(window).scrollTop()),10)) {
                mainmenu.addClass("fixed");
            } else {
                mainmenu.removeClass("fixed");
            }
        });
        jQuery(".cs-main-navigation-wrapper.cs-menu-is-sticky").wrap("<div class='cs-menu-parent'></div>").attr("rel", jQuery(".cs-main-navigation-wrapper.cs-menu-is-sticky").offset().top).parent().height($(".cs-main-navigation-wrapper.cs-menu-is-sticky").height());

        // Ingredients
        jQuery(".cs-ingredients-check-list ul li").on("click", function(){
            jQuery(this).toggleClass("active");
        });

        // Custom choosen form elements
        jQuery(".chosen-select").chosen({width: "100%"});

        // Lightbox - image
        jQuery(".cs-lightbox-image").magnificPopup({
            type: "image",
            mainClass: "mfp-with-zoom",
            zoom: {
                enabled: true,
                duration: 300,
                easing: "ease-in-out",
            }
        });

        // Lightbox - gallery
        jQuery(".cs-lightbox-gallery").each(function () {
            jQuery(this).magnificPopup({
                delegate: "a",
                type: "image",
                gallery: {
                    enabled: true
                },
                mainClass: "mfp-with-zoom",
                zoom: {
                    enabled: true,
                    duration: 300,
                    easing: "ease-in-out",
                }
            });
        });

        // Lightbox - iframe
        jQuery(".cs-lightbox-iframe").magnificPopup({
            type: "iframe",
            mainClass: "mfp-with-zoom",
            zoom: {
                enabled: true,
                duration: 300,
                easing: "ease-in-out",
            }
        });

        // Countdown
        jQuery(".cs-countdown").countdown("2016/12/12", function(event) {
            jQuery(this).html(event.strftime(''
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%-m</div>'
                    + '<div class="cs-countdown-label">month%!m</div>'
                + '</div>'
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%-w</div>'
                    + '<div class="cs-countdown-label">week%!w</div>'
                + '</div>'
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%-d</div>'
                    + '<div class="cs-countdown-label">day%!d</div>'
                + '</div>'
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%H</div>'
                    + '<div class="cs-countdown-label">hour%!H</div>'
                + '</div>'
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%M</div>'
                    + '<div class="cs-countdown-label">minute%!M</div>'
                + '</div>'
                + '<div class="cs-countdown-block">'
                    + '<div class="cs-countdown-number">%S</div>'
                    + '<div class="cs-countdown-label">seconds</div>'
                + '</div>'
            ));
        });

        // Init post gallery
        var mySwiper = new Swiper (".cs-post-gallery", {
            grabCursor: true,
            prevButton: ".cs-swiper-button-prev",
            nextButton: ".cs-swiper-button-next",
            autoHeight: true
        });

        // Init home slider
        var mySwiper = new Swiper (".cs-home-slider .swiper-container", {
            grabCursor: true,
            prevButton: ".cs-swiper-button-prev",
            nextButton: ".cs-swiper-button-next",
            autoHeight: true
        });

        // Init home regular slider
        var mySwiper = new Swiper (".cs-home-regular-slider", {
            grabCursor: true,
            prevButton: ".cs-swiper-button-prev",
            nextButton: ".cs-swiper-button-next",
            autoHeight: true
        });

        // Init recipe gallery slider
        var galleryTop = new Swiper(".cs-recipe-gallery", {
            grabCursor: true,
            prevButton: ".cs-swiper-button-prev",
            nextButton: ".cs-swiper-button-next",
            autoHeight: true
        });
        var galleryThumbs = new Swiper(".cs-recipe-thumbs", {
            centeredSlides: true,
            slidesPerView: 'auto',
            slideToClickedSlide: true
        });
        galleryTop.params.control = galleryThumbs;
        galleryThumbs.params.control = galleryTop;

        // Filter by tags (isotope)
        var $grid = $(".cs-isotope-recipes").isotope({
            itemSelector: ".cs-isotope-recipe",
            layoutMode: "fitRows"
        });

        // Filter items on button click
        $(".cs-isotope-filter").on( "click", "button", function() {
            var filterValue = $(this).attr("data-filter");
            $grid.isotope({ filter: filterValue });
            $(".cs-isotope-filter button").removeClass("current");
            $(this).addClass("current");
        });

    });
})(jQuery);

