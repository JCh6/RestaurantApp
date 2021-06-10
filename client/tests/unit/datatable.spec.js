import Vue from 'vue';
import Vuetify from 'vuetify';
import DataTable from '@/components/DataTable.vue';
import { createLocalVue, mount } from '@vue/test-utils';

Vue.use(Vuetify);

describe('DataTable.vue', () => {
    const localVue = createLocalVue()
    let vuetify;

    beforeEach(() => {
        vuetify = new Vuetify()
    })

    const mountFunction = options => {
        return mount(DataTable, {
            localVue,
            vuetify,
            ...options,
        })
    }

    it('should create a table and match the headers', () => {
        const wrapper = mountFunction({
            propsData: {
                header: [
                    { text: "ID", value: "id" },
                    { text: "Name", value: "name" },
                    { text: "Age", value: "age" },
                ],
                data: [{}],
                handleClick: () => {}
            },
        })

        expect(wrapper.html()).toContain("ID");
        expect(wrapper.html()).toContain("Name");
        expect(wrapper.html()).toContain("Age");
    })
})