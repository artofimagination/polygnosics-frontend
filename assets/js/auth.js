
!function($) {
  "use strict";
  if ($("#signup").length) {
    $("#signup-loader").get(0).style.display = "none"
    $("#signup-dialog").get(0).style.display = "block"

    
    var SweetAlertSignup = function() {};
    SweetAlertSignup.prototype.init = function() {
      $('#signup').submit(function(){
        $("#signup-loader").get(0).style.display = "block"
        $("#signup-dialog").get(0).style.display = "none"
        $.ajax({
          type: 'POST',
          url: $(this).attr('action'),
          data: $(this).serialize(),
          dataType: 'json',
          error: function(data) {
            $("#signup-loader").get(0).style.display = "none"
            $("#signup-dialog").get(0).style.display = "block"
            var alertType  = "error"
            if (data.responseText == "Registration successful")
            {
              alertType = "success"
            }

            swal({
              title: "Registration", 
              text: data.responseText, 
              type: alertType
              },
              function(){
                window.location.href = "/index";
              }) 
          }
        })
        return false;        
      });
    },
    //init
    $.SweetAlertSignup = new SweetAlertSignup, $.SweetAlertSignup.Constructor = SweetAlertSignup
  }

  if ($("#signin").length) {
    var SweetAlertSignin = function() {};
    SweetAlertSignin.prototype.init = function() {
      $('#signin').submit(function(){
        $.ajax({
          type: 'POST',
          url: $(this).attr('action'),
          data: $(this).serialize(),
          dataType: 'json',
          error: function(data) {
            if (data.status === 202)
            {
              swal({
                title: "Login", 
                text: data.responseText, 
                type: "error"
                },
                function(){
                  window.location.href = "/auth_login";
                })
              return 
            }
            window.location.href = "/user-main";
          }
        })
        return false;        
      });
    },
    //init
    $.SweetAlertSignin = new SweetAlertSignin, $.SweetAlertSignin.Constructor = SweetAlertSignin
  }
}(window.jQuery),

//initializing 
function($) {
  "use strict";
  if ($("#signup").length) {
    $.SweetAlertSignup.init()
  }
  if ($("#signin").length) {
    $.SweetAlertSignin.init()
  }
}(window.jQuery);