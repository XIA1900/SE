describe('login', function(){
    this.beforeEach(()=>{
            cy.visit('http://localhost:8000/user/login')  
        })

    it('correctLogin', function(){
    cy.get('.ant-pro-form-login-main').get('#username').type('kirby')
    cy.get('.ant-pro-form-login-main').get('#password').type('007')
    cy.get('button:contains("Login")').click()
    cy.url().should('include', '/homepage')
    })
    it('incorrectLogin', function(){
    cy.get('.ant-pro-form-login-main').get('#username').type('admin')
    cy.get('.ant-pro-form-login-main').get('#password').type('ant')
    cy.get('button:contains("Login")').click()
    cy.url().should('include', '/user/login') 
    })
    it('forwardPostpage', function(){
        cy.contains('Apple').click()
        cy.url().should('include', '/group/content?22')
        })
})
