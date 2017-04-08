var Solitaire = {};
$(function () {

    Solitaire.URL = {
        Game: "game"
    };

    Solitaire.Tempalates = {
        /**
         * @return {string}
         */
        InfoHtml: function () {
            var infoHtml = "";
            infoHtml +=
                "<p>Move #: " +
                (Solitaire.CurrentMove + 1) +
                "</p>";
            infoHtml =
                "<div class='group info'>" +
                infoHtml +
                "</div>";
            return infoHtml
        },
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

            var infoHtml = Solitaire.Tempalates.InfoHtml() + "<br/>";
            var foundationsHtml = Solitaire.Tempalates.FoundationsHtml(game.Foundations);
            var deckHtml = Solitaire.Tempalates.DeckHtml(game.Deck);
            var pilesHtml = Solitaire.Tempalates.PilesHtml(game.Piles);

            $ele.html(infoHtml + foundationsHtml + deckHtml + pilesHtml);
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

    /** @type {FullGame} Solitaire.CurrentFullGame **/
    Solitaire.CurrentFullGame = {};
    /** @type {Game} Solitaire.CurrentGame **/
    Solitaire.CurrentGame = {};
    /** @type {int} Solitaire.CurrentMove **/
    Solitaire.CurrentMove = -1;

    Solitaire.FullGame = {
        /**
         * @param {jQuery} $ele
         * @param {FullGame} fullGame
         */
        SetGame: function ($ele, fullGame) {
            Solitaire.CurrentFullGame = Solitaire.FullGame.Copy(fullGame);
            /**
             * @type {Game}
             */
            Solitaire.CurrentGame = {
                Foundations: [
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null}
                ],
                Deck: fullGame.Deck,
                Piles: fullGame.Piles
            };
            Solitaire.CurrentMove = -1;
            Solitaire.Tempalates.Game($ele, Solitaire.CurrentGame);
        },
        /**
         * @param {FullGame} fullGame
         */
        Copy: function (fullGame) {
            return JSON.parse(JSON.stringify(fullGame))
        },
        /**
         * @param {jQuery} $ele
         * @param {int} moveNumber
         */
        RenderState: function ($ele, moveNumber) {
            var fullGame = Solitaire.FullGame.Copy(Solitaire.CurrentFullGame);
            Solitaire.CurrentGame = {
                Foundations: [
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null},
                    {Cards: null, Suit: null}
                ],
                Deck: fullGame.Deck,
                Piles: fullGame.Piles
            };
            console.log("Move #" + moveNumber + ":");
            if (moveNumber >= 0) {
                console.log(Solitaire.CurrentFullGame.Moves[moveNumber]);
            }
            for (var i = 0; i <= moveNumber; i++) {
                Solitaire.FullGame.DoMove(i);
            }
            Solitaire.CurrentMove = moveNumber;
            Solitaire.Tempalates.Game($ele, Solitaire.CurrentGame);
        },
        /**
         * @param {int} moveNum
         */
        DoMove: function (moveNum) {
            var move = Solitaire.CurrentFullGame.Moves[moveNum];
            if (move.SourcePileId === PileDeck) {
                /** From Deck */
                Solitaire.FullGame.Moves.FromDeck(move);
            } else {
                /** From Pile */
                Solitaire.FullGame.Moves.FromPile(move);
            }
        },
        Moves: {
            /**
             * @param {Move} move
             */
            FromDeck: function (move) {
                var card;
                if (move.TargetPileId === PileDeck) {
                    /** Deck flip */
                    Solitaire.CurrentGame.Deck.Position++;
                    if (Solitaire.CurrentGame.Deck.Position > Solitaire.CurrentGame.Deck.Cards.length) {
                        Solitaire.CurrentGame.Deck.Position = 0;
                    }
                } else {
                    /** To Pile or Foundation */
                    card = Solitaire.CurrentGame.Deck.Cards.splice(Solitaire.CurrentGame.Deck.Position - 1, 1)[0];
                    Solitaire.CurrentGame.Deck.Position--;
                    if (move.TargetPileId === PileFoundation) {
                        /** To Foundation */
                        Solitaire.FullGame.Moves.ToFoundation(card);
                    } else {
                        /** To Pile */
                        Solitaire.CurrentGame.Piles[move.TargetPileId].StackCards.push(card);
                    }
                }
            },
            /**
             * @param {Move} move
             */
            FromPile: function (move) {
                var card, cards, i;
                cards = Solitaire.CurrentGame.Piles[move.SourcePileId].StackCards.splice(move.SourcePileIndex, Solitaire.CurrentGame.Piles[move.SourcePileId].StackCards.length);
                if (move.TargetPileId === PileFoundation) {
                    /** To Foundation */
                    Solitaire.FullGame.Moves.ToFoundation(cards[0]);
                } else {
                    /** To Pile */
                    for (i = 0; i < cards.length; i++) {
                        Solitaire.CurrentGame.Piles[move.TargetPileId].StackCards.push(cards[i]);
                    }
                }
                var notSubCardMove = move.SourcePileIndex === 0;
                var hasBaseCards = Solitaire.CurrentGame.Piles[move.SourcePileId].BaseCards && Solitaire.CurrentGame.Piles[move.SourcePileId].BaseCards.length > 0;
                if (notSubCardMove && hasBaseCards) {
                    card = Solitaire.CurrentGame.Piles[move.SourcePileId].BaseCards.splice(0, 1)[0];
                    Solitaire.CurrentGame.Piles[move.SourcePileId].StackCards.push(card);
                }
            },
            /**
             * @param {Card} card
             */
            ToFoundation: function(card) {
                for (var i = 0; i < Solitaire.CurrentGame.Foundations.length; i++) {
                    if (Solitaire.CurrentGame.Foundations[i].Suit === null || Solitaire.CurrentGame.Foundations[i].Suit === card.Suit) {
                        if (Solitaire.CurrentGame.Foundations[i].Cards === null) {
                            Solitaire.CurrentGame.Foundations[i].Cards = [];
                        }
                        Solitaire.CurrentGame.Foundations[i].Cards.push(card);
                        Solitaire.CurrentGame.Foundations[i].Suit = card.Suit;
                        break;
                    }
                }
            }
        }
    };

    Solitaire.Form = {
        /**
         * @param {jQuery} $game
         * @param {jQuery} $gotoMoveForm
         */
        GotoMove: function ($game, $gotoMoveForm) {
            $gotoMoveForm.submit(function(e) {
                e.preventDefault();
                var val = $gotoMoveForm.find("[type=text]").val();
                if (val == 0) {
                    val = 0;
                }
                val--;
                Solitaire.FullGame.RenderState($game, val);
            });
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
                    /** @type {FullGame} fullGame */
                    var fullGame = JSON.parse(data);
                } catch (e) {
                    console.log(e);
                    return;
                }
                console.log("Full game:");
                console.log(fullGame);
                Solitaire.FullGame.SetGame($ele, fullGame);
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
        var moveNumber = Solitaire.CurrentMove + 1;
        if (moveNumber >= Solitaire.CurrentFullGame.Moves.length) {
            moveNumber = 0;
        }
        Solitaire.FullGame.RenderState($ele, moveNumber);
    };

    /**
     * @param {jQuery} $ele
     */
    Solitaire.PrevMove = function ($ele) {
        var moveNumber = Solitaire.CurrentMove - 1;
        if (moveNumber < -1) {
            moveNumber = Solitaire.CurrentFullGame.Moves.length - 1;
        }
        Solitaire.FullGame.RenderState($ele, moveNumber);
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
 *   SourcePileIndex: int
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
