var Solitaire = {};
$(function () {

    Solitaire.URL = {
        Game: "game",
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

            for (i = 0; i < game.Foundations.length; i++) {
                var foundation = game.Foundations[i];
                tmpHtml = "";
                if (!foundation.Cards) {
                    tmpHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                } else {
                    for (j = 0; j < foundation.Cards.length; j++) {
                        card = foundation.Cards[j];
                        var number = card.Number;
                        switch (number) {
                            case 1:
                                card = "A";
                                break;
                            case 11:
                                card = "J";
                                break;
                            case 12:
                                card = "Q";
                                break;
                            case 13:
                                card = "k";
                                break;
                        }
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
                if (game.Deck.Cards) {
                    deckHtml += Solitaire.Tempalates.Snippets.CardFlipped();
                } else {
                    deckHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                }
            } else {
                var startCard = game.Deck.Position - 3;
                if (startCard < 0) {
                    startCard = 0;
                }
                if (startCard > 0) {
                    deckHtml += Solitaire.Tempalates.Snippets.CardFlipped();
                }
                for (i = startCard; i < game.Deck.Position; i++) {
                    card = game.Deck.Cards[i];
                    deckHtml = Solitaire.Tempalates.Snippets.Card(card);
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
                if (pile.BaseCards) {
                    for (j = 0; j < pile.BaseCards.length; j++) {
                        tmpHtml += Solitaire.Tempalates.Snippets.CardFlipped();
                    }
                }
                if (pile.StackCards) {
                    for (j = 0; j < pile.StackCards.length; j++) {
                        card = pile.StackCards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.Card(card);
                    }
                }
                if (!pile.StackCards && !pile.BaseCards) {
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
                return "<div class='card " + card.Suit + "'>" +
                    "<span>" + card + "</span>" +
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
            CardFlipped: function () {
                return "<div class='card flipped'></div>";
            }
        }
    };

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
        })
    };

});

/**
 * @typedef {{
 *   Foundations: []Foundation
 *   Deck: Deck
 *   Piles: []Pile
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
