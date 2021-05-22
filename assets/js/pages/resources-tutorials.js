
!function($) {
  "use strict";

  function switchElements() {
    if ($("#video").is(':checked')) {
      $("#avatar_video").show();
      $("#avatar_image").hide();
    } else {
      $("#avatar_video").hide();
      $("#avatar_image").show();
    }
  }

  switchElements()
  $(document).ready(function() {
    $("#radio-list input").change("click", function() {
      switchElements();
    });

    $("#create-new-item").click(function(){
      var form
      if ($("#tutorials-form").length) {
        form = $("#tutorials-form")
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
      var formData = new FormData(form[0]);
      $.ajax({
        type: 'POST',
        url: form.attr('action'),
        data: formData,
        enctype: form.attr('enctype'),
        contentType: false,
        processData: false,
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
          window.location.href = "/resources/tutorials";
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
