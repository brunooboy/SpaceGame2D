package game

import (
	"game2d/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfW,
		Y: 500,
	}

	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	// atualizar a lógica do jogador aqui
	speed := 6.0

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.position.X -= speed
		if p.position.X < 0 {
			p.position.X = 0
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.position.X += speed
		bounds := p.image.Bounds()
		playerWidth := float64(bounds.Dx())
		if p.position.X+playerWidth > screenWidth {
			p.position.X = screenWidth - playerWidth
		}
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.IsReady() {
		p.laserLoadingTimer.Reset()

		// posição do jogador na tela
		bounds := p.image.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfW,
			p.position.Y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	// Posição do jogador na tela
	op.GeoM.Translate(p.position.X, p.position.Y)

	// Desenhar o jogador na tela
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()))
}
