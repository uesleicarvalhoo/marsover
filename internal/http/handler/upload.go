package handler

import (
	"net/http"

	"github.com/uesleicarvalhoo/marsrover/internal/http/formatter"
	"github.com/uesleicarvalhoo/marsrover/internal/http/parser"
	"github.com/uesleicarvalhoo/marsrover/internal/http/utils"
	"github.com/uesleicarvalhoo/marsrover/orchestrator"
)

// UploadMission godoc
// @Summary     Envia arquivo de missão
// @Description Recebe um arquivo de texto com as instruções dos rovers e retorna as posições finais em formato plain-text.
// @Tags        Missions
// @Accept      multipart/form-data
// @Produce     plain
// @Param       file formData file true "Arquivo de missão"
// @Success     200 {string} string "Exemplo: 1 3 N\n5 1 E"
// @Failure     400 {object} utils.APIError "Requisição malformada (form inválido)"
// @Failure     405 {object} utils.APIError "Método não permitido"
// @Failure     422 {object} utils.APIError "Erro de domínio ao processar o arquivo"
// @Failure     500 {object} utils.APIError "Erro interno inesperado"
// @Router      /missions [post]
func MissionFromFile(svc orchestrator.MissionUseCase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			utils.WriteError(w, utils.New(http.StatusMethodNotAllowed, "method not allowed", nil))
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, 2<<20)
		if err := r.ParseMultipartForm(2 << 20); err != nil {
			utils.WriteError(w, utils.New(http.StatusBadRequest, "invalid form: "+err.Error(), nil))
			return
		}

		file, _, err := r.FormFile("file")
		if err != nil {
			utils.WriteError(w, utils.New(http.StatusBadRequest, "missing file: "+err.Error(), nil))
			return
		}
		defer file.Close()

		params, err := parser.ParseMission(r.Context(), file)
		if err != nil {
			utils.WriteError(w, utils.New(http.StatusUnprocessableEntity, err.Error(), nil))
			return
		}

		results, err := svc.Execute(r.Context(), params)
		if err != nil {
			utils.WriteError(w, utils.New(http.StatusUnprocessableEntity, err.Error(), nil))
			return
		}

		res := formatter.FormatResultsPlain(results)

		utils.WriteText(w, res, http.StatusOK)
	}
}
