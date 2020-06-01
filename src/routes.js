import GameList from './components/GameList.vue'
import GameSingle from './components/Game'
import RuleList from './components/RuleList'
import RuleSingle from './components/Rule'

export const routes = [
    {
        path: "/",
        component: GameList
    },
    {
        path: "/game/:id",
        name: "Game",
        component: GameSingle,
        children: [
            {
                path: "rules",
                name: "Game Rules",
                components: {
                    page: RuleList
                },
                meta: {
                    showModal: false
                }
            },
            {
                path: "rules/:ruleId",
                name: "Game Rule",
                components: {
                    page: RuleList,
                    rule: RuleSingle
                },
                meta: {
                    showModal: true
                }
            }
        ]
    }
]

