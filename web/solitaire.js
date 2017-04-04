var Solitaire = {};
$(function () {

    Solitaire.URL = {
        Game: "game",
        Reset: "reset"
    };

    Solitaire.Tempalates = {
        /**
         * @param {jQuery} $ele
         * @param {Game} game
         */
        Game: function ($ele, game) {
            var i, j, card;
            var foundationsHtml = "";
            var deckHtml = "";
            var pilesHtml = "";
            var tmpHtml = "";

            console.log("Moves: " + game.Moves + ", Position: " + game.Deck.Position + ", Size: " + game.Deck.Cards.length);

            for (i = 0; i < game.Foundations.length; i++) {
                var foundation = game.Foundations[i];
                tmpHtml = "";
                if (foundation.Cards === null) {
                    tmpHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                } else {
                    for (j = 0; j < foundation.Cards.length; j++) {
                        card = foundation.Cards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.Card(card);
                    }
                }
                foundationsHtml +=
                    "<div class='pile'>" +
                    tmpHtml +
                    "</div>";
            }
            foundationsHtml =
                "<div class='group foundation'>" +
                "<h4>Foundations</h4>" +
                foundationsHtml +
                "</div>";

            if (game.Deck.Position === 0) {
                if (game.Deck.Cards.length > 0) {
                    card = game.Deck.Cards[0];
                    deckHtml += Solitaire.Tempalates.Snippets.CardFlipped(card);
                } else {
                    deckHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                }
            } else {
                var startCard = game.Deck.Position - 3;
                if (startCard < 0) {
                    startCard = 0;
                }
                if (game.Deck.Position > 0) {
                    if (game.Deck.Position === game.Deck.Cards.length) {
                        deckHtml += Solitaire.Tempalates.Snippets.CardEmptyDeck();
                    } else {
                        card = game.Deck.Cards[game.Deck.Position];
                        deckHtml += Solitaire.Tempalates.Snippets.CardFlippedDeck(card);
                    }
                }
                for (i = startCard; i < game.Deck.Position; i++) {
                    card = game.Deck.Cards[i];
                    deckHtml += Solitaire.Tempalates.Snippets.Card(card);
                }
            }
            deckHtml =
                "<div class='group deck'>" +
                "<h4>Deck</h4>" +
                "<div class='pile'>" +
                deckHtml +
                "</div>" +
                "</div>";

            for (i = 0; i < game.Piles.length; i++) {
                var pile = game.Piles[i];
                tmpHtml = "";
                if (pile.BaseCards.length > 0) {
                    for (j = pile.BaseCards.length - 1; j >= 0; j--) {
                        card = pile.BaseCards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.CardFlipped(card);
                    }
                }
                if (pile.StackCards.length > 0) {
                    for (j = 0; j < pile.StackCards.length; j++) {
                        card = pile.StackCards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.Card(card);
                    }
                }
                if (pile.StackCards.length === 0 && pile.BaseCards.length === 0) {
                    tmpHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                }
                pilesHtml +=
                    "<div class='pile'>" +
                    tmpHtml +
                    "</div>";
            }
            pilesHtml =
                "<div class='group piles'>" +
                "<h4>Piles</h4>" +
                pilesHtml +
                "</div>";

            $ele.html(foundationsHtml + deckHtml + pilesHtml);
        },
        Snippets: {
            /**
             * @param {Card} card
             * @return {string}
             */
            Card: function (card) {
                var number = GetCardLetter(card.Number);
                return "<div class='card " + card.Suit + "'>" +
                    "<span>" + number + "</span>" +
                    "<div class='suit'></div>" +
                    "</div>";
            },
            /**
             * @return {string}
             */
            CardEmpty: function () {
                return "<div class='card empty'></div>";
            },
            /**
             * @return {string}
             */
            CardEmptyDeck: function () {
                return "<div class='card empty-deck'></div>";
            },
            /**
             * @param {Card} card
             * @return {string}
             */
            CardFlippedDeck: function (card) {
                var number = GetCardLetter(card.Number);
                return "<div class='card flipped-deck " + card.Suit + "'>" +
                    "<span>" + number + "</span>" +
                    "<div class='suit'></div>" +
                    "</div>";
            },
            /**
             * @param {Card} card
             * @return {string}
             */
            CardFlipped: function (card) {
                var number = GetCardLetter(card.Number);
                return "<div class='card flipped " + card.Suit + "'>" +
                    "<span>" + number + "</span>" +
                    "<div class='suit'></div>" +
                    "</div>";
            },
            /**
             * @return {string}
             */
            CardDeck: function () {
                return "<div class='card deck'></div>";
            }
        }
    };

    /**
     * @param {int} cardNumber
     * @return {string}
     */
    function GetCardLetter(cardNumber) {
        var number = "" + cardNumber;
        switch (cardNumber) {
            case 1:
                number = "A";
                break;
            case 11:
                number = "J";
                break;
            case 12:
                number = "Q";
                break;
            case 13:
                number = "K";
                break;
        }
        return number;
    }

    /**
     * @param {jQuery} $ele
     */
    Solitaire.GetGame = function ($ele) {
        $.ajax({
            url: Solitaire.URL.Game,
            success: function (data) {
                try {
                    var game = JSON.parse(data);
                } catch (e) {
                    console.log(e);
                    return;
                }
                Solitaire.Tempalates.Game($ele, game);
            },
            error: function (err) {
                console.log(err);
            }
        });
    };

    /**
     * @param {jQuery} $ele
     */
    Solitaire.ResetGame = function($ele) {
        $.ajax({
            url: Solitaire.URL.Reset,
            success: function (data) {
                try {
                    var game = JSON.parse(data);
                } catch (e) {
                    console.log(e);
                    return;
                }
                for (var i = 0; i < game.Deck.Cards.length; i++) {
                    var card = game.Deck.Cards[i];
                    console.log(GetCardLetter(card.Number) + " " + card.Suit);
                }
                Solitaire.Tempalates.Game($ele, game);
            },
            error: function (err) {
                console.log(err);
            }
        });
    };

});

/**
 * @typedef {{
 *   Foundations: []Foundation
 *   Deck: Deck
 *   Piles: []Pile
 *   Moves: int
 * }} Game
 */

/**
 * @typedef {{
 *   Cards: []Card
 *   Suit: string
 * }} Foundation
 */

/**
 * @typedef {{
 *   Cards: []Card
 *   Position: int
 * }} Deck
 */

/**
 * @typedef {{
 *   BaseCards: []Card
 *   StackCards: []Card
 * }} Pile
 */

/**
 * @typedef {{
 *   Number: int
 *   Suit: string
 * }} Card
 */
