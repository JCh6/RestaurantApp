import Vue from 'vue';
import Vuetify from 'vuetify';
import Footer from '@/components/Footer.vue';
import { createLocalVue, mount } from '@vue/test-utils';

Vue.use(Vuetify);

describe('Footer.vue', () => {
    const localVue = createLocalVue()
    let vuetify;

    beforeEach(() => {
        vuetify = new Vuetify()
    })

    const mountFunction = options => {
        return mount(Footer, {
            localVue,
            vuetify,
            ...options,
        })
    }

    it('manipulates state | setting a footer link', async () => {
        const wrapper = mountFunction({});

        await wrapper.setData(
            { images: [{
                link: "https://github.com/JCh6/RestaurantApp",
                name: "github.svg",
                target: "_blank",
            }] 
        })
        
        expect(wrapper.html()).toContain("https://github.com/JCh6/RestaurantApp");
    })
})