!function($) {
  "use strict";
    //bootstrap WYSIHTML5 - text editor
    var editor = $('.textarea').wysihtml5({
      toolbar: {
        "font-styles": true, // Font styling, e.g. h1, h2, etc.
        "emphasis": true, // Italics, bold, etc.
        "lists": true, // (Un)ordered lists, e.g. Bullets, Numbers.
        "html": true, // Button which allows you to edit the generated HTML.
        "link": true, // Button to insert a link.
        "image": true, // Button to insert an image.
        "color": false, // Button to change color of font
        "blockquote": true, // Blockquote
        "size": "sm" // options are xs, sm, lg
      },
      parserRules: {
        classes: {
          'wysiwyg-color-silver' : 1,
          'wysiwyg-color-gray' : 1,
          'wysiwyg-color-white' : 1,
          'wysiwyg-color-maroon' : 1,
          'wysiwyg-color-red' : 1,
          'wysiwyg-color-purple' : 1,
          'wysiwyg-color-fuchsia' : 1,
          'wysiwyg-color-green' : 1,
          'wysiwyg-color-lime' : 1,
          'wysiwyg-color-olive' : 1,
          'wysiwyg-color-yellow' : 1,
          'wysiwyg-color-navy' : 1,
          'wysiwyg-color-blue' : 1,
          'wysiwyg-color-teal' : 1,
          'wysiwyg-color-aqua' : 1,
          'wysiwyg-color-orange' : 1
        },
        tags: {
          'b':  {},
          'i':  {},
          'strong': {},
          'em': {},
          'p': {},
          'br': {},
          'ol': {},
          'ul': {},
          'li': {},
          'h1': {},
          'h2': {},
          'h3': {},
          'h4': {},
          'h5': {},
          'h6': {},
          'blockquote': {},
          'u': 1,
          'img': {
            'check_attributes': {
              'width': 'numbers',
              'alt': 'alt',
              'src': 'url',
              'height': 'numbers'
            }
          },
          'a':  {
            'check_attributes': {
              'href': 'url'
            },
            'set_attributes': {
              'target': '_blank',
              'rel': 'nofollow'
            }
          },
          'span': 1,
          'div': 1,
          'small': 1,
          'code': 1,
          'pre': 1
        }
      }
    });

    var innerHTML = $("<li>\n" +
                   "  <div id='insertVideoModal' class='bootstrap-wysihtml5-insert-video-modal modal fade' data-wysihtml5-dialog='insertVideo' style='display: none;' aria-hidden='true'>\n" +
                   "    <div class='modal-dialog '>\n" +
                   "      <div class='modal-content'>\n" +
                   "        <div class='modal-header'>\n" +
                   "          <script>console.log('insertVideoModal')</script>\n" +
                   "          <h3>Insert Video</h3>\n" +
                   "          <a class='close' data-dismiss='modal'>×</a>\n" +
                   "        </div>\n" +
                   "        <div class='modal-body'>\n" +
                   "          <div class='form-group'>\n" +
                   "            <input value='http://' class='bootstrap-wysihtml5-insert-link-url form-control' data-wysihtml5-dialog-field='href'>\n" +
                   "          </div>\n" +
                   "          <div class='checkbox'>\n" +
                   "            <input type='checkbox' class='bootstrap-wysihtml5-insert-link-target' checked=''>\n" +
                   "            <label>Open link in new window</label>\n" +
                   "          </div>\n" +
                   "        </div>\n" +
                   "        <div class='modal-footer'>\n" +
                   "          <a class='btn btn-default' data-dismiss='modal' data-wysihtml5-dialog-action='cancel' href='#'>Cancel</a>\n" +
                   "          <a href='#' class='btn btn-primary' data-dismiss='modal' data-wysihtml5-dialog-action='save'>Insert link</a>\n" +
                   "        </div>\n" +
                   "      </div>\n" +
                   "    </div>\n" +
                   "  </div>\n" +
                   "  <a id='insertVideo' class='btn btn-sm btn-default' data-wysihtml5-command='insertVideo' title='Insert video' tabindex='-1' href='javascript:;' unselectable='on'>\n" +
                   "    <span class='glyphicon glyphicon-film'></span>\n" +
                   "  </a>\n" +
                   "</li>\n")
    $('.wysihtml5-toolbar').append(innerHTML);

    // TODO Issue#153: Need to finish adding insert video button for wysihtml editor
    // For the time being just enable edit mode and paste iframe manually
    // <div class="embed-responsive embed-responsive-16by9">
    //    <iframe width="400" height="200" src="https://www.youtube.com/embed/fn3KWM1kuAw" frameborder="0" allowfullscreen></iframe>
    // </div>
    wysihtml5.commands.insertVideo = {
      exec: function(composer, command, value) {
      },
      state: function(composer, command, className) {
        return wysihtml5.commands.formatInline.state(composer, command, "A");
      }
    }

    var dialog = new wysihtml5.toolbar.Dialog(
      document.querySelector("[data-wysihtml5-command='insertVideo']"),
      document.querySelector("[data-wysihtml5-dialog='insertVideo']")
    );

    dialog.observe("save", function(attributes) {

    });   

    var SweetAlertCreate = function() {};
    SweetAlertCreate.prototype.init = function() {
      $('#create-new-item').submit(function(){
        $.ajax({
          type: 'POST',
          url: $(this).attr('action'),
          data: $(this).serialize(),
          dataType: 'json',
          error: function(data) {
            if (data.status === 202)
            {
              swal({
                title: "Add new item", 
                text: data.responseText, 
                type: "error"
                },
                function(){
                  window.location.href = "/user-main";
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
    $.SweetAlertCreate = new SweetAlertCreate, $.SweetAlertCreate.Constructor = SweetAlertCreate
}(window.jQuery),

//initializing 
function($) {
  "use strict";

  $.SweetAlertCreate.init()
}(window.jQuery);

