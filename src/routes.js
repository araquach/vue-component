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
        name: "game",
        component: GameSingle,
        children: [
            {
                path: "rules",
                name: "game-rules",
                component: RuleList,
                meta: {
                    showModal: false
                }
            },
            {
                path: "rules/:ruleId",
                name: "game-rule",
                components: {
                    default: RuleList,
                    rule: RuleSingle
                },
                meta: {
                    showModal: true
                }
            }
        ]
    }
]

