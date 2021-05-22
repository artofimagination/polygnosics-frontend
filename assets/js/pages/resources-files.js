
!function($) {
  "use strict";
  var count = $('#files-group .radio-list').length
  function switchElements(index) {
    var linkName = "#repo_link_"+index
    var fileName = "#upload_file_"+index
    if ($("#link_"+index).is(':checked')) {
      $(linkName).show();
      $(fileName).hide();
    } else {
      $(linkName).hide();
      $(fileName).show();
    }
  }

  for (var i = 0; i<count; i++){
    switchElements(i)
  }
  
  $(document).ready(function() {
    for (var i = 0; i<count; i++){
      $("#radio-list_" + i + " input").change("click", function() {
        var res = $(this).attr('id').split("_");
        switchElements(res[1]);
      });
    }
    
    $("#add-more").click(function(){
      var innerHTML =
      "<label class='mt-40'>Reference text :<span class='danger'>*</span></label>\n" +
      "<input type='text' class='form-control required' id='ref_name_" + count + "' name='ref_name_" + count + "'>\n" +
      "<div id='radio-list_" + count + "' class='radio-list'>\n" +
        "<label class='radio-inline p-0 mr-10'>\n" +
          "<div class='radio radio-info'>\n" +
            "<input type='radio' name='type_" + count + "' id='link_" + count + "' value='link' checked>\n" +
            "<label for='link_" + count + "'>Repo Link</label>\n" +
          "</div>\n" +
        "</label>\n" +
        "<label class='radio-inline'>\n" +
          "<div class='radio radio-info'>\n" +
            "<input type='radio' name='type_" + count + "' id='file_" + count + "' value='file'>\n" +
            "<label for='file_" + count + "'>File</label>\n" +
          "</div>\n" +
        "</label>\n" +
      "</div>\n" +
      "<input type='url' class='form-control' id='repo_link_" + count + "' name='repo_link_" + count + "'>\n" +
      "<input type='file' class='btn btn-primary form-control' id='upload_file_" + count + "' name='upload_file_" + count + "'>\n";
      $('#upload_file_' + (count - 1)).after(innerHTML);
      switchElements(count);
      $("#radio-list_" + count + " input").on("click", function() {
        var res = $(this).attr('id').split("_");
        switchElements(res[1]);
      });
      if (count < 50) {
        count++;
      }
      return false;        
    });

    $("#create-new-item").click(function(){
      var form
      if ($("#files-form").length) {
        form = $("#files-form")
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
      formData.append("count", count);
      for(var pair of formData.entries()) {
        if (pair[0].includes("repo_link_")) {
          var res = pair[0].split("_");
          var link = "type_"+res[2];
          if (formData.get(link) !== "file") {
            if (pair[1].includes("github")) {
              formData.set(link, "github")
            } else if (pair[1].includes("gitlab")) {
              formData.set(link, "gitlab")
            } else if (pair[1].includes("bitbucket")) {
              formData.set(link, "bitbucket")
            } else{
              formData.set(link, "link")
            }
          }
        }
      }

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
          window.location.href = "/resources/files";
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
