var Solitaire = {};
$(function () {

    Solitaire.URL = {
        Game: "game",
        FullGame: "full-game",
        Reset: "reset"
    };

    Solitaire.Tempalates = {
        /**
         * @param {[Foundation]} foundations
         * @return {string}
         */
        FoundationsHtml: function (foundations) {
            var foundationsHtml = "";
            var tmpHtml = "";
            var card, i;
            for (i = 0; i < foundations.length; i++) {
                var foundation = foundations[i];
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
            return foundationsHtml;
        },
        /**
         * @param {Deck} deck
         * @return {string}
         */
        DeckHtml: function (deck) {
            var deckHtml = "";
            var card, i;
            if (deck.Position === 0) {
                if (deck.Cards.length > 0) {
                    card = deck.Cards[0];
                    deckHtml += Solitaire.Tempalates.Snippets.CardFlipped(card);
                } else {
                    deckHtml += Solitaire.Tempalates.Snippets.CardEmpty();
                }
            } else {
                var startCard = deck.Position - 3;
                if (startCard < 0) {
                    startCard = 0;
                }
                if (deck.Position > 0) {
                    if (deck.Position === deck.Cards.length) {
                        deckHtml += Solitaire.Tempalates.Snippets.CardEmptyDeck();
                    } else {
                        card = deck.Cards[deck.Position];
                        deckHtml += Solitaire.Tempalates.Snippets.CardFlippedDeck(card);
                    }
                }
                for (i = startCard; i < deck.Position; i++) {
                    card = deck.Cards[i];
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
            return deckHtml;
        },
        /**
         * @param {[Pile]} piles
         * @return {string}
         */
        PilesHtml: function (piles) {
            var pilesHtml = "";
            var tmpHtml = "";
            var card, i, j;
            for (i = 0; i < piles.length; i++) {
                var pile = piles[i];
                tmpHtml = "";
                var hasBaseCards = pile.BaseCards && pile.BaseCards.length > 0;
                var hasStackCards = pile.StackCards && pile.StackCards.length > 0;
                if (hasBaseCards) {
                    for (j = pile.BaseCards.length - 1; j >= 0; j--) {
                        card = pile.BaseCards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.CardFlipped(card);
                    }
                }
                if (hasStackCards) {
                    for (j = 0; j < pile.StackCards.length; j++) {
                        card = pile.StackCards[j];
                        tmpHtml += Solitaire.Tempalates.Snippets.Card(card);
                    }
                }
                if (!hasBaseCards && !hasStackCards) {
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
            return pilesHtml;
        },
        /**
         * @param {jQuery} $ele
         * @param {Game} game
         */
        Game: function ($ele, game) {
            //console.log("Moves: " + game.Moves + ", Position: " + game.Deck.Position + ", Size: " + game.Deck.Cards.length);

            var foundationsHtml = Solitaire.Tempalates.FoundationsHtml(game.Foundations);
            var deckHtml = Solitaire.Tempalates.DeckHtml(game.Deck);
            var pilesHtml = Solitaire.Tempalates.PilesHtml(game.Piles);

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

    var PileDeck = 7;
    var PileFoundation = 8;

    /** @type {FullGame} CurrentFullGame **/
    var CurrentFullGame;
    /** @type {Game} CurrentGame **/
    var CurrentGame;
    /** @type {int} CurrentMove **/
    var CurrentMove;
    Solitaire.FullGame = {
        /**
         * @param {jQuery} $ele
         * @param {Game} game
         */
        SetGame: function ($ele, game) {
            CurrentFullGame = game;
            //console.log(game);
            /**
             * @type {Game}
             */
            CurrentGame = {
                Foundations: [
                    { Cards: null, Suit: null },
                    { Cards: null, Suit: null },
                    { Cards: null, Suit: null },
                    { Cards: null, Suit: null }
                ],
                Deck: game.Deck,
                Piles: game.Piles
            };
            CurrentMove = 0;
            Solitaire.Tempalates.Game($ele, CurrentGame);
        },
        /**
         * @param {jQuery} $ele
         */
        NextMove: function($ele) {
            var move = CurrentFullGame.Moves[CurrentMove];
            console.log(move);
            var card, cards, i;
            if (move.SourcePileId === PileDeck) {
                if (move.TargetPileId === PileDeck) {
                    CurrentGame.Deck.Position++;
                    if (CurrentGame.Deck.Position > CurrentGame.Deck.Cards.length) {
                        CurrentGame.Deck.Position = 0;
                    }
                } else {
                    card = CurrentGame.Deck.Cards.splice(CurrentGame.Deck.Position - 1, 1)[0];
                    CurrentGame.Deck.Position--;
                    if (move.TargetPileId === PileFoundation) {
                        for (i = 0; i < CurrentGame.Foundations.length; i++) {
                            if (CurrentGame.Foundations[i].Suit === null || CurrentGame.Foundations[i].Suit === card.Suit) {
                                if (CurrentGame.Foundations[i].Cards === null) {
                                    CurrentGame.Foundations[i].Cards = [];
                                }
                                CurrentGame.Foundations[i].Cards.push(card);
                                CurrentGame.Foundations[i].Suit = card.Suit;
                                break;
                            }
                        }
                    } else {
                        CurrentGame.Piles[move.TargetPileId].StackCards.push(card);
                    }
                }
            } else {
                cards = CurrentGame.Piles[move.SourcePileId].StackCards.splice(0, CurrentGame.Piles[move.SourcePileId].StackCards.length);
                if (move.TargetPileId === PileFoundation) {
                    for (i = 0; i < CurrentGame.Foundations.length; i++) {
                        if (CurrentGame.Foundations[i].Suit === null || CurrentGame.Foundations[i].Suit === cards[0].Suit) {
                            if (CurrentGame.Foundations[i].Cards === null) {
                                CurrentGame.Foundations[i].Cards = [];
                            }
                            CurrentGame.Foundations[i].Cards.push(cards[0]);
                            CurrentGame.Foundations[i].Suit = cards[0].Suit;
                            break;
                        }
                    }
                } else {
                    for (i = 0; i < cards.length; i++) {
                        CurrentGame.Piles[move.TargetPileId].StackCards.push(cards[i]);
                    }
                }
                if (CurrentGame.Piles[move.SourcePileId].BaseCards && CurrentGame.Piles[move.SourcePileId].BaseCards.length > 0) {
                    card = CurrentGame.Piles[move.SourcePileId].BaseCards.splice(0, 1)[0];
                    CurrentGame.Piles[move.SourcePileId].StackCards.push(card);
                }
            }
            Solitaire.Tempalates.Game($ele, CurrentGame);
            CurrentMove++;
        }
    };

    /**
     * @param {jQuery} $ele
     */
    Solitaire.GetGame = function ($ele) {
        $.ajax({
            url: Solitaire.URL.FullGame,
            success: function (data) {
                try {
                    var game = JSON.parse(data);
                } catch (e) {
                    console.log(e);
                    return;
                }
                console.log(game);
                Solitaire.FullGame.SetGame($ele, game);
            },
            error: function (err) {
                console.log(err);
            }
        });
    };

    /**
     * @param {jQuery} $ele
     */
    Solitaire.NextMove = function ($ele) {
        Solitaire.FullGame.NextMove($ele);
    };

    /**
     * @param {jQuery} $ele
     */
    Solitaire.ResetGame = function ($ele) {
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
 *   Deck: Deck
 *   Piles: []Pile
 *   Moves: []Move
 *   Won: bool
 * }} FullGame
 */

/**
 * @typedef {{
 *   SourceCard: Card
 *   TargetCard: Card
 *   SourcePileId: int
 *   TargetPileId: int
 * }} Move
 */

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
