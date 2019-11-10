package handles

import "errors"

func insertDinheiro(user string, valor float64) {
	trans := new(Transacao)
	trans.ID = queryGetUseID(user)
	trans.Descricao = "deposito"
	trans.Tipo = 1
	trans.Valor = valor
	insertTransacao(*trans)
}

func insertCompra(userCompra, codigo string, valor float64, qtd int) error {

	lances := queryGetCarteiraUser(codigo, valor)
	for i := 0; i < qtd; i++ {

		transVenda := new(Transacao)
		transVenda.ID = lances[i].IDUser
		transVenda.Tipo = 2
		transVenda.Valor = valor
		transVenda.Descricao = "venda acao:" + codigo

		if transVenda.ID == -1 {
			err := errors.New("usuario de venda não encontrado")
			return err
		}

		transCompra := new(Transacao)
		transCompra.ID = queryGetUseID(userCompra)
		transCompra.Descricao = "compra acao" + codigo
		transCompra.Tipo = 3
		transCompra.Valor = -valor

		if transCompra.ID == -1 {
			err := errors.New("usuario de compra não encontrado")
			return err
		}
		saldo := queryGetSaldo(transCompra.ID)

		if saldo < valor {
			err := errors.New("sem saldo")
			return err
		}

		carteira := new(Carteira)
		carteira.Codigo = codigo
		carteira.ID = transCompra.ID
		carteira.Lance = 0
		carteira.Venda = false
		insertCarteira(*carteira)
		removeCarteira(lances[i].IDCarteira)
		insertTransacao(*transVenda)
		insertTransacao(*transCompra)
	}

	return nil
}
