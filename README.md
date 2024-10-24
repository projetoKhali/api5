<br id="topo">

![Read me Khali](https://github.com/user-attachments/assets/94aecab2-e751-4ab4-a2a8-1a6589b4eb01)

<p align="center">
    <a href="#sobre">Sobre</a> |
    <a href="#backlogproduto">Backlog do produto</a>  |
    <a href="#backlogsprint">Backlog da sprint </a> |
    <a href="#entrega">Entrega das Sprints</a>  |
    <a href="#burndown">Burndown</a>  |
    <a href="#gerenciamento">Gerenciamento do projeto</a>  |
    <a href="#documentacoes">Documentações</a>  |
    <a href="#prototipo">Protótipo</a>     |
    <a href="#tecnologias">Tecnologias</a>  |
    <a href="#equipe">Equipe</a>        
</p>

<span id="sobre">

## Desafio

O desafio apresentado pela empresa <a href="https://www.pro4tech.com.br/"> Pro4tech </a> consiste em aprimorar a eficiência e eficácia no processo de recrutamento e seleção, buscando otimizar a coleta, visualização e análise dos dados. A principal dificuldade está na fragmentação dessas informações, que precisam ser centralizadas para facilitar decisões estratégicas, melhorar a alocação de recursos e alinhar as estratégias de contratação aos objetivos organizacionais.

## Proposta

Nossa proposta consiste em um dashboard interativo que centraliza informações essenciais sobre o processo de recrutamento. Ele fornecerá insights valiosos e permitirá uma análise detalhada e personalizada das etapas de seleção.

- Funcionalidades Principais:
    - Métricas úteis: Exibição de tempo médio de contratação, status das vagas, número de processos em andamento, e outras informações cruciais para a análise de desempenho.
    - Filtros personalizados: Permitirão segmentar as informações com base em critérios relevantes, como processo, tipo de vaga, entre outros.
    - Geração de relatórios: Opção de extração de relatórios detalhados para análise offline, facilitando o acompanhamento e a documentação dos processos.
    - Sistema de notificações: Envio de alertas automáticos aos recrutadores quando um processo seletivo estiver próximo de expirar, garantindo que ações sejam tomadas no tempo adequado.
    - Gerenciamento de usuários: Sistema de gerenciamento de perfis, com diferentes níveis de acesso, permitindo que usuários tenham permissões ajustadas conforme suas funções.

Essa solução oferece uma visão clara e abrangente do processo seletivo, permitindo uma tomada de decisão estratégica, melhor alocação de recursos e agilidade operacional.

<span id="backlogproduto">

## Backlog do produto

<details>

<summary>Confira nosso Backlog</summary>
    
| Épico  | Story | Prioridade | Critérios de aceite | Código do Requisito | Sprint |
|-----------------|------------------------------|------------------------------------------------------|------------|----------------|----------------|
|  Dashboard processo seletivo  | Eu como Gestor de RH quero um dashboard para acompanhar as métricas dos processos seletivos e gerir seu progresso e prazos   | Muito Importante |  - O gráfico de barras deve exibir o tempo médio de contratação em comparação com os meses <br> - O gráfico de pizza deve mostrar o status das vagas  (Aberta, Fechada, Em Análise) <br> - Os cards devem exibir as principais métricas do processo seletivo  <br> - Os filtros poderão ser usados para uma melhor para uma melhor personalização do dashboard| RF-01 | 1 |
| Dashboard processo seletivo  | Eu como Analista de RH quero um dashboard para acompanhar as métricas dos processos seletivos e gerir seu progresso e prazos  | Muito Importante |  - O gráfico de barras mostra o tempo médio de contratação em comparação com os meses <br> - O gráfico de pizza mostra o status das vagas  (Aberta, Fechada, Em Análise) <br> - Os cards devem exibir as principais métricas do processo seletivo  <br> - Os filtros poderão ser usados para uma melhor para uma melhor personalização do dashboard | RF-01 | 1 |
|  Dashboard processo seletivo  | Eu como Gestor do setor quero um dashboard para acompanhar as métricas dos processos seletivos e gerir seu progresso e prazos   | Muito Importante | - O gráfico de barras mostra o tempo médio de contratação em comparação com os meses <br> - O gráfico de pizza mostra o status das vagas  (Aberta, Fechada, Em Análise) <br> - Os cards devem exibir as principais métricas do processo seletivo  <br> - Os filtros poderão ser usados para uma melhor para uma melhor personalização do dashboard| RF-01 | 1 |
| Dashboard vaga               | Eu como Gestor de RH quero visualizar em dashboard as métricas das vagas de emprego para ter um controle das contratações | Importante |  - Dashboard exibe métricas atualizadas em tempo real <br> - Filtros para uma melhor personalização do dashboard <br> - Dados atualizados em tempo real  | 	RF-03 | 2 |
|  Dashboard vaga               | Eu como Analista de RH quero visualizar em dashboard as métricas das vagas de emprego para ter um controle das contratações  | Importante |  - Visualização clara das vagas em aberto <br> - Filtros para uma melhor personalização do dashboard <br> - Dados atualizados em tempo real | 	RF-03 | 2 |
 | Dashboard vaga               | Eu como Gestor do setor quero visualizar em dashboard as métricas das vagas de emprego para ter um controle das contratações   | Importante |  - Visualização clara das vagas em aberto da área do gestor <br> - Filtros para uma melhor personalização do dashboard <br> - Dados atualizados em tempo real | 	RF-03 | 2 |
| Extração de relatório        | Eu como Gestor de RH quero extrair relatórios com os dados do processo seletivo em formatos como PDF e Excel para apresentações e análises offline| Importante | - Relatório exportado para PDF e Excel com sucesso <br> - Dados exportados corretamente e filtros aplicados <br> - Formatação do relatório adequada para ambos os formatos| RF-04 | 3 |
| Extração de relatório        | Eu como Analista de RH quero extrair relatórios com os dados do processo seletivo em formatos como PDF e Excel para apresentações e análises offline   | Importante | - Relatório exportado para PDF e Excel com sucesso <br> - Dados exportados corretamente e filtros aplicados <br> - Formatação do relatório adequada para ambos os formatos | RF-04 | 3 | 
| Extração de relatório        | Eu como Gestor do setor quero extrair relatórios com os dados do processo seletivo em formatos como PDF e Excel para apresentações e análises offline  | Importante | - Relatório sobre meu setor exportado para PDF e Excel com sucesso <br> - Dados exportados corretamente e filtros aplicados <br> - Formatação do relatório adequada para ambos os formatos | RF-04 | 3 |
| Extração de relatório        | Como Gestor de RH, quero poder filtrar e personalizar os relatórios extraídos para que eu possa adaptar os dados às minhas necessidades e facilitar a gestão.                     | Desejável  |  - Funcionalidade de personalização de relatórios disponível <br> - Filtros customizáveis aplicados corretamente <br> - Relatórios ajustados com base nas preferências | 	RF-02 | 3 | 
|  Extração de relatório        | Eu como Analista de RH quero poder filtrar e personalizar os relatórios extraídos para que eu possa adaptar os dados às minhas necessidades e facilitar a gestão.                 | Desejável  |   - Funcionalidade de personalização de relatórios disponível <br> - Filtros customizáveis aplicados corretamente <br> - Relatórios ajustados com base nas preferências | RF-02 | 3 |
|  Extração de relatório        | Eu como Gestor do setor quero poder filtrar e personalizar os relatórios extraídos para que eu possa adaptar os dados às minhas necessidades e facilitar a gestão.                | Desejável  |  - Filtros aplicáveis e customizáveis <br> - Relatórios gerados de acordo com as especificações do setor <br> - Exportação com formatação correta | RF-02 | 3 |
|  Notificações                 | Eu como Gestor de RH quero ser notificado sempre que um indicador chave for atingido para eu me manter atualizado   | Importante | - Notificações geradas automaticamente ao atingir indicadores chave <br> - Frequência e formato das notificações customizáveis | 	RF-06 |
|  Notificações                 | Eu como Analista de RH quero ser notificado sempre que um indicador chave for atingido para eu me manter atualizado    | Importante |  - Notificações geradas automaticamente ao atingir indicadores chave <br> - Frequência e formato das notificações customizáveis | 	RF-06 |
|  Notificações                 | Eu como Gestor do setor quero ser notificado sempre que um indicador chave for atingido para eu me manter atualizado  | Desejável  |  - Notificações geradas automaticamente ao atingir indicadores chave <br> - Frequência e formato das notificações customizáveis | 	RF-06 |
|  Indicadores Chaves           | Eu como Gestor de RH quero poder selecionar indicadores chaves para acompanhar seu progresso  | Importante | 	- Seleção e personalização de indicadores chave disponíveis <br> - Progresso visível em tempo real <br> - Atualização automática dos dados dos indicadores | RF-06 |
|  Indicadores Chaves | Eu como Analista de RH quero poder selecionar indicadores chaves para acompanhar seu progresso    | Importante |  	- Seleção e personalização de indicadores chave disponíveis <br> - Progresso visível em tempo real <br> - Atualização automática dos dados dos indicadores | RF-06 |
| Indicadores Chaves | Eu como Gestor do setor quero poder selecionar indicadores chaves para acompanhar seu progresso | Desejável   | 	- Seleção e personalização de indicadores chave disponíveis <br> - Progresso visível em tempo real <br> - Atualização automática dos dados dos indicadores | RF-06 |
|  Gestão de Usuários           | Eu como Gestor de RH quero poder cadastrar usuários para que eles possam acessar o sistema | Importante | - Formulário de cadastro funcional </br> - Validação de campos obrigatórios </br> - Usuário pode acessar o sistema após o cadastro  | RF-01 |
| Gestão de Usuários           | Eu como Gestor de RH quero poder editar o cadastro de usuários para manter o sistema atualizado | Importante | - Usuário existente editado com sucesso <br> - Permissões modificadas são refletidas no acesso <br> - Confirmação de edição bem-sucedida | RF-01 |
| Gestão de Usuários           | Eu como Gestor de RH quero poder remover usuários para que eles não tenham mais acesso ao sistema | Desejável  |  - Usuário removido com sucesso <br> - Usuário perde imediatamente o acesso <br> - Mensagem de confirmação da remoção | RF-01 |
| Cadastro de grupos de acesso | Eu como Gestor de RH quero poder criar e gerenciar grupos de usuários com permissões específicas para que cada usuário tenha acesso somente ao que foi pré definido | Importante |  - Criação de grupos de acesso com permissões configuráveis <br > - Usuários atribuídos a grupos e acessos refletidos corretamente <br> - Confirmação de criação e edição de grupos | RF-01 |

</details>

<details>
<summary>Requisitos do projeto </summary>

| Código  | Tipo           | Descrição                                                                                                        | Prioridade |
|---------|----------------|------------------------------------------------------------------------------------------------------------------|------------|
| RF-01   | Funcional      | Desenvolver um dashboard interativo para visualizar métricas de recrutamento e seleção em tempo real.             | Alta       |
| RF-02   | Funcional      | Permitir a personalização dos relatórios com filtros por departamento, tipo de vaga, e outras categorias.         | Alta       |
| RF-03   | Funcional      | Fornecer análises detalhadas sobre o desempenho do processo de recrutamento, incluindo o tempo médio de contratação e taxas de retenção. | Média      |
| RF-04   | Funcional      | Gerar relatórios automáticos e exportáveis em formatos como PDF e Excel para apresentações e análises offline.     | Média      |
| RF-05   | Funcional      | Facilitar o compartilhamento de relatórios e dashboards entre equipes e departamentos com controle de acesso e permissões. | Baixa      |
| RF-06   | Funcional      | Permitir a seleção de indicadores chave de desempenho e a configuração de alertas automáticos para quando esses indicadores atingirem determinados níveis. | Baixa      |
| RNF-01  | Não Funcional  | Manual do Usuário detalhado para orientar os usuários na utilização do dashboard e na geração de relatórios.        | Alta       |
| RNF-02  | Não Funcional  | Guia de instalação para a configuração inicial da plataforma, incluindo requisitos de hardware e software.          | Média      |
| RNF-03  | Não Funcional  | Modelagem de Banco de Dados para armazenamento eficiente e seguro dos dados de recrutamento e seleção.              | Alta       |
| RNF-04  | Não Funcional  | Implementação de protocolos de segurança para proteger dados sensíveis dos candidatos e da empresa, conforme as melhores práticas de segurança da informação. | Alta       |

</details>


<span id="backlogsprint">
    
## Backlog da sprint

<details open>
<summary>Sprint 1 </summary>


| **Épico** | **Funcionalidade** | **Pontuação** |
| :-------- | :-------- | :-----------: |
| Dashboard Processo Seletivo          | Eu como usuário do sistema quero um dashboard para acompanhar as métricas dos processos seletivos em tempo real e gerir seu progresso e prazos       |      9       |
| Dashboard Processo Seletivo          | Eu como usuário do sistema quero personalizar as métricas disponíveis no dashboard para filtrar informações relevantes de acordo com as minhas necessidades | 8 |

</details>
</br>
<details open>

<summary> Sprint 2 </summary>

| **Épico** | **Funcionalidade**                  | **Pontuação** |
| :-------- | :---------------------------------- | :-----------: |
| Dashboard Vagas  |  Eu como usuário do sistema quero visualizar no dashboard as métricas das vagas de emprego para ter um controle das contratações   |       8       |
| Dashboard Vagas  |  Eu como usuário do sistema quero poder personalizar as métricas do dashboard por recrutador, status do processo e status da vaga  |       7       |

</details>
</br>

<details open> 
    
<summary> Sprint 3 </summary>

| **Épico** | **Funcionalidade**                  | **Pontuação** |
| :-------- | :---------------------------------- | :-----------: |
| Extração de relatório   |  Eu como usuário do sistema quero extrair relatórios com os dados do processo seletivo em formatos como PDF e Excel para apresentações e análises offline   |               |
| Extração de relatório   |  Eu como usuário do sistema quero poder filtrar e personalizar os relatórios para que eu possa extrair somente os dados necessários                         |               |

</details>
</br>


<span id="entrega">


## Entrega das sprints

<details>

<summary>Sprint 1 - Dashboard de Processo Seletivo </summary>

### De 09/09 a 29/09
#
### Dashboard de Processo Seletivo

- Acompanhamento de Processos Seletivos
- Desenvolvimento de um dashboard interativo com gráficos e métricas que mostram o progresso dos processos seletivos em tempo real. As informações exibidas no dashboard incluem:

    - Cards que exibem as principais métricas do processos seletivos.
    - Gráfico de barras que exibe o tempo médio de contratação em comparação com os meses.
    - Gráfico de pizza que exibe os status das vagas.

### Wireframe

- Criação do wireframe das telas do dashboard para aprovação e visualização prévia por parte do usuário, permitindo feedback e ajustes antes do desenvolvimento final.

### Modelagem de Dados

- Modelagem do banco de dados com as tabelas necessárias para armazenar informações sobre os processos seletivos, vagas, candidatos e status.

### Personalização das Métricas

- Funcionalidade que permite ao usuário personalizar as métricas exibidas, com a opção de aplicar filtros por processo seletivo, vagas e período, facilitando uma análise mais direcionada conforme as necessidades de gestão.

### Video demonstrando o funcionamento da aplicação


https://github.com/user-attachments/assets/5b3d1251-f504-41c7-8af6-ef9a4df80c7d



</details>
</br>

<details>

<summary>Sprint 2 - Dashboard de vagas </summary>

### De 30/09 a 20/10
#
### Dashboard de Vagas

- Acompanhamento de vagas
- Desenvolvimento de um dashboard interativo com gráficos e métricas que mostram o progresso das vagas em tempo real. As informações exibidas no dashboard que foram incluidas nesta sprint são:

    - Tabela contendo dados importantes sobre as vagas.

### Wireframe

- Modificações no wireframe das telas do dashboard para aprovação e visualização prévia por parte do usuário, permitindo feedback e ajustes antes do desenvolvimento final.

### Personalização das Métricas

- Funcionalidade que permite ao usuário personalizar as métricas exibidas, foram adicionados mais filtros para que a personalização seja ainda mais precisa.

</details>
</br>

<span id="burndown">

## Burndown

<details>
<summary> Sprint 1 </summary>
    
![Burndown Sprint 1](https://github.com/user-attachments/assets/5e6c02a5-1f58-48db-82af-56f5b042da2c)


</details>
</br>

<details>
<summary> Sprint 2</summary>

![Burndown sprint 2](https://github.com/user-attachments/assets/bc124a35-13a7-4951-b5fb-a4da6220051f)


</details>

<span id="gerenciamento">

## Gerenciamento do projeto 

<a href="https://github.com/orgs/projetoKhali/projects/21/views/5">Github Projects</a>

<span id="documentacoes">

## Documentações 

>Confira nossas documentações: https://github.com/projetoKhali/docs-project.git
<br>

>Confira nosso manual do usuário: https://sites.google.com/view/projetokhali/in%C3%ADcio

<span id="prototipo">

## Protótipo

<a href="https://www.figma.com/proto/UAmnY0AL8mKeqezOnPq2Kh/API-5-Prot%C3%B3tipo?node-id=18-2&node-type=canvas&t=o2y0WxLgkEQBoBy6-1&scaling=min-zoom&content-scaling=fixed&page-id=0%3A1">Figma</a>

<span id="tecnologias">

</br>

## Tecnologias

<h3 style="color: #C1FF72  "> Front-end </h3>

<img height= 80 src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/react/react-original.svg" />

> React Native
          
<h3 style="color: #C1FF72  "> Back-end </h3>

<img height= 100 src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original-wordmark.svg" />

> Go         
          

#


<span id="equipe">

![Read me Khali grupo](https://github.com/user-attachments/assets/b79f54f3-8ae6-4380-af1c-49a6143440b1)


<table style="width:100%; border-collapse: collapse;">
    <tr style="background-color: #620874; color: #06EF47;">
        <th style="text-align: center; text-align: center; padding: 10px;">Função</th>
        <th style="text-align: center; text-align: center; padding: 10px;">Nome</th>
        <th style="text-align: center; text-align: center; padding: 10px;">LinkedIn & GitHub</th>
    </tr>
    <tr>
        <td style="text-align: center; text-align: center; padding: 10px;">Scrum Master</td>
        <td style="text-align: center; text-align: center; padding: 10px;">Jhonatan Lopes</td>
        <td style="text-align: center; text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/jhonatan-o-lopes/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/JhonatanLop"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; text-align: center; padding: 10px;">Product Owner</td>
        <td style="text-align: center; padding: 10px;">Nicole Souza</td>
        <td style="text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/nicolem-souza/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/NicSouza"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; text-align: center; padding: 10px;">Developer</td>
        <td style="text-align: center; text-align: center; padding: 10px;">Paulo Granthon</td>
        <td style="text-align: center; text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/paulo-granthon/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/paulo-granthon"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; padding: 10px;">Developer</td>
        <td style="text-align: center; padding: 10px;">Markus Gomes</td>
        <td style="text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/markus-gomes-013b76250"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/markusgomes"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; padding: 10px;">Developer</td>
        <td style="text-align: center; text-align: center; padding: 10px;">Marcos Vinícius Malaquias</td>
        <td style="text-align: center; text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/marcos-malaquias/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/Incivius"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; padding: 10px;">Developer</td>
        <td style="text-align: center; padding: 10px;">Tânia Cruz</td>
        <td style="text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/tânia-cruz-30ab5812a/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/taniacruzz"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
    <tr>
        <td style="text-align: center; padding: 10px;">Developer</td>
        <td style="text-align: center; padding: 10px;">Ludmila Chagas</td>
        <td style="text-align: center; padding: 10px;">
            <a href="https://www.linkedin.com/in/ludmila-chagas-273548187/"><img src="https://img.shields.io/badge/-Linkedin-blue?style=flat-square&logo=Linkedin&logoColor=white" alt="LinkedIn"></a>
            <a href="https://github.com/ludmila-chagas"><img src="https://img.shields.io/badge/-GitHub-111217?style=flat-square&logo=github&logoColor=white" alt="GitHub"></a>
        </td>
    </tr>
</table>



→ [Voltar ao topo](#topo)
