var Solitaire = {};
$(function() {

    Solitaire.URL = {
        Game: "game",
    };

    Solitaire.GetGame = function() {
        $.ajax({
            url: Solitaire.URL.Game,
            success: function(data) {
                try {
                    var parsed = JSON.parse(data);
                } catch(e) {
                    console.log(e);
                    return;
                }
                console.log(parsed);
            },
            error: function(err) {
                console.log(err);
            }
        })
    };

});
