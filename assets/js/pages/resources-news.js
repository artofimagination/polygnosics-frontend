
!function($) {
  "use strict";

  
  $(document).ready(function() {
    $("#create-new-item").click(function(){
      var form
      if ($("#news-form").length) {
        form = $("#news-form")
      }else{
        swal({
          title: "Add new item", 
          text: "Invalid form", 
          type: "error"
          },
          function(){
            window.location.href = "/user-main";
          })
      }
      
      $.ajax({
        type: 'POST',
        url: form.attr('action'),
        data: form.serialize(),
        error: function(data) {
          swal({
            title: "Add new item", 
            text: data.responseText, 
            type: "error"
            },
            function(){
              window.location.href = form.attr('action');
            })
        },
        success: function(data) {
          window.location.href = "/resources/news";
        }
      })
      return false;        
    }); 
  });

}(window.jQuery),

//initializing 
function($) {
  "use strict";
}(window.jQuery);
