// gpt-3.5-turbo-instruct, prompted with PGN, defeeats Stockfish level 4 (1700?)
// Date: Sep 18, 2023
// Link: https://lichess.org/D39lnanQ
//
// Procedure: I started a game on LiChess against Stockfish Level 4, prompted GPT with the PGN
// preamble followed by "1. ", got the completion of Nf3, played that in LiChess, added in
// Stockfish's reply of Nf6 followed by "\n2. ", and repeated. After move 29, GPT kept responding
// with 1-0, so I added the prompt about Anand explaining how he would have forced mate, at which
// point GPT forced the win. 
//
// h/t @GrantSlatton on X/Twitter for the tip
//
// On Sep 18, 2023, @GrantSlatton wrote (https://x.com/GrantSlatton/status/1703913578036904431):
//
// > The new GPT model, gpt-3.5-turbo-instruct, can play chess around 1800 Elo.
// > I had previously reported that GPT cannot play chess, but it appears this was just the RLHF'd
//   chat models. The pure completion model succeeds.
// > The new model readily beats Stockfish Level 4 (1700) and still loses respectably to Level 5
//   (2000). Never attempted illegal moves. Used clever opening sacrifice, and incredibly cheeky pawn
//   & king checkmate, allowing the opponent to uselessly promote.
// > https://lichess.org/K6Q0Lqda
// > I used this PGN style prompt to mimic a grandmaster game.
// > The highlighting is a bit wrong. GPT made all its own moves, I input Stockfish moves manually.
// > h/t to @zswitten for this prompt style
// > [OpenAI Playground screenshot showing PGN game in the prompt]
//
// https://lichess.org/forum/general-chess-discussion/stockfish-level-and-its-rating :
// OGeffert wrote:
//
// "I think Stockfish level 4 - announced with a rating of 1700"
//

package main

import (
	"context"
	"fmt"

	"codegene.com/openai/openaikey"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient(openaikey.MustGet())
	resp, err := client.CreateCompletion(
		context.Background(),
		openai.CompletionRequest{
			Model:       "gpt-3.5-turbo-instruct",
                  Temperature: 0.2,
		  Prompt: `[Event "Shamkir Chess"]
[White "Anand, Viswanathan"]
[Black "Topalov, Veselin"]
[Result "1-0"]
[WhiteElo "2779"]
[BlackElo "2740"]

1. Nf3 Nf6
2. c4 e6
3. Nc3 d5
4. d4 Be7
5. Bg5 O-O
6. e3 Nbd7
7. Qc2 h6
8. Bh4 g5
9. Bg3 g4
10. Ne5 c6
11. O-O-O dxc4
12. Bxc4 Qe8
13. h3 b5
14. Bd3 b4
15. Ne4 Nxe4
16. Bxe4 b3
17. axb3 Nxe5
18. Bxe5 f6
19. Bxc6 Bd7
20. Bxd7 fxe5
21. Bxe8 Rfxe8
22. hxg4 Kg7
23. Kb1 exd4
24. Rx4 Rf8
25. Rd7 Rae8
26. Qc7 Kg6
27. Rxe7 Rxe7
28. Qxe7 Rf7
29. Qxe6+ Kg5
30. 1-0

After the game, Anand explained to a group of novices how he could have forced mate by playing

31. Qxh6+ Kxg4
32. Qh5#
`,
})

	if err != nil {
		fmt.Printf("Completion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Text)
}
