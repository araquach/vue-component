<template>
    <div>
        <h1 class="title is-3">People List</h1>
        <ul>
            <li v-for="person in people" :key="person.id">
                <router-link :to="{ name: 'person', params: { id: person.id }}" :people="people">
                    {{ person.firstName }}
                </router-link>
            </li>
        </ul>
        <b-modal :active.sync="showModal"
                 trap-focus
                 aria-role="dialog"
                 aria-modal>
            <router-view name="person"/>
        </b-modal>
    </div>
</template>

<script>
    import people from "../data/people";

    export default {
        data() {
            return {
                people
            }
        },
        watch: {
            "$route.meta"({showModal}) {
                this.showModal = showModal;
            }
        }
    }
</script>