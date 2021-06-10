import Vue from 'vue';
import Vuetify from 'vuetify';
import BackButton from '@/components/BackButton.vue';
import { createLocalVue, mount } from '@vue/test-utils';

Vue.use(Vuetify);

describe('BackButton.vue', () => {
    const localVue = createLocalVue()
    let vuetify;

    beforeEach(() => {
        vuetify = new Vuetify()
    })

    const mountFunction = options => {
        return mount(BackButton, {
            localVue,
            vuetify,
            ...options,
        })
    }

    it('should have a custom msg, icon, route, color and match snapshot', () => {
        const wrapper = mountFunction({
            propsData: {
                msg: "Back to Calendar",
                icon: "Calendar",
                color: "primary",
            },
        })

        expect(wrapper.html()).toMatchSnapshot();
    })
})