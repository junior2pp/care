$(".menu-bar").on("click", function() {
  $(".paginas").toggleClass("abrir")
})

$("li a").click(function() {

  $("a").removeClass("activo")

  $(this).toggleClass("activo")

})

$(".submenu").click(function() {
  $(this).children("ul").slideToggle();
})


$("ul").click(function(p) {
  p.stopPropagation();
})
