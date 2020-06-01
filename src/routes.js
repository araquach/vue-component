import Home from "./components/Home"
import People from "./components/People"
import Person from "./components/Person"

export const routes = [
    {
        path: "",
        name: "home",
        component: Home
    },
    {
        path: "/people",
        name: "people",
        component: People,
        children: [
            {
                path: "/person/:id",
                name: "person",
                components: {
                    default: People,
                    person: Person
                },
                meta: {
                    showModal: true
                },
                props: true
            }
        ]
    },
]