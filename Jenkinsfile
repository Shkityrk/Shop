def deployBackup() {
    echo "DEPLOY:FAIL"
    echo "Попытка отката"

    if (fileExists("/var/Shop/backup_repo")) {
        echo "Восстановление из резервной копии"
        sh 'rm -rf ./*'
        sh 'cp -r /var/Shop/backup_repo/* ./'
        echo "Запускаем старую версию"
        sh 'docker compose -f .\prod.docker-compose.yml up --build -d'
        echo "Сбой новой версии, выполнен откат"

    } else {
        echo "Бэкап не найден, откат невозможен!"
        error("Бэкап не найден")
    }
}


def createBackup() {
    echo "Создание резервной копии текущего репозитория"

    if (fileExists("/var/Shop/backup_repo")) {
        echo "Удаляем предыдущий бэкап"
        sh 'rm -rf /var/Shop/backup_repo'
    }

    echo "Создаём новый бэкап"
    sh 'mkdir -p /var/Shop/backup_repo'
    sh 'cp -r . /var/Shop/backup_repo'
    echo "Репозиторий скопирован в /var/Shop/backup_repo"
}



pipeline {
    agent any

    triggers {
        githubPush()
    }

    environment {
        NETWORK_NAME = "shop_network"
    }

    stages {
        stage('Docker Access Info') {
            steps {
                script {
                    sh '''
                    echo "Пользователь:"
                    whoami

                    echo "Группы:"
                    groups || id

                    echo "Текущая директория:"
                    pwd
                    echo "Содержимое директории:"
                    ls -la

                    echo "Docker доступ:"
                    which docker
                    docker version || echo "Не удалось получить версию Docker"
                    '''

                }
            }
        }


        stage('Check Network Existence') {
            steps {
                script {
                    def networkExists = sh(
                        script: "docker network inspect ${NETWORK_NAME} > /dev/null 2>&1 && echo 'exists' || echo 'not_exists'",
                        returnStdout: true
                    ).trim()

                    if (networkExists == 'exists') {
                        echo "Сеть ${NETWORK_NAME} существует"
                        env.NETWORK_EXISTS = "true"
                    } else {
                        echo "Сеть ${NETWORK_NAME} не существует"
                        env.NETWORK_EXISTS = "false"
                    }
                }
            }
        }


        stage('Create backup of repo') {
            steps {
                script {
                    createBackup()
                }
            }
        }


        stage('Stop and Remove Containers if exist') {
            when {
                expression { return env.NETWORK_EXISTS == "true" }
            }
            steps {
                script {
                    echo "docker compose -f .\prod.docker-compose.yml down"


                    def runningContainers = sh(
                        script: "docker ps -q --filter network=${NETWORK_NAME}",
                        returnStdout: true
                    ).trim()

                    if (runningContainers) {
                        echo "Останавливаем контейнеры: ${runningContainers}"
                        sh "docker stop ${runningContainers}"
                        sh "docker system prune -a"
                    } else {
                        echo "Нет запущенных контейнеров в сети"
                    }


                    def allContainers = sh(
                        script: "docker ps -aq --filter network=${NETWORK_NAME}",
                        returnStdout: true
                    ).trim()

                    if (allContainers) {
                        echo "Удаляем контейнеры: ${allContainers}"
                        sh "docker rm -f ${allContainers}"
                    } else {
                        echo "Нет контейнеров для удаления в сети"
                    }
                }
            }
        }


        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/main']],
                    userRemoteConfigs: [[
                        url: 'https://github.com/Shkityrk/Shop.git',
                    ]]
                ])
            }
        }

        stage('Prepare Environment Directory') {
            steps {
                script {
                    sh '''
                    if [ -d "/var/Shop/environment" ]; then

                        if [ -d "./environment" ]; then
                            echo "Удаляем старую ./environment"
                            rm -rf ./environment
                        fi
                        cp -r /var/Shop/environment ./environment
                    else
                        echo "ERROR:Папка /var/Shop/environment не найдена!"
                        exit 1
                    fi
                    '''
                }
            }
        }

        stage('Deploy ') {
            steps {
                script {

                    echo "docker compose -f .\prod.docker-compose.yml up --build -d"
                    try {
                        sh 'docker compose -f .\prod.docker-compose.yml up --build -d'
                        echo "DEPLOY:OK"
                    } catch (Exception e) {
                        echo "Ошибка при деплое: ${e.getMessage()}"
                        deployBackup()
                    }
                }
            }
        }

    stage('Pause 30 seconds') {
        steps {
            echo 'Ждём 30 секунд перед healthcheck…'
            sleep(time: 30, unit: 'SECONDS')
        }
    }

        stage('Healthcheck') {
    steps {
        script {
            def healthy = false

            for (int i = 1; i <= 5; i++) {
                echo "Healthcheck попытка #${i}"
                // --connect-timeout 1 и --max-time 1 гарантируют,
                // что curl не уйдёт дольше секунды
                def code = sh(
                    script: """
                      curl -s -o /dev/null -w '%{http_code}' http://79.137.197.216/auth/healthcheck
                    """,
                    returnStdout: true
                ).trim()

                if (code == '200') {
                    echo "Попытка #${i}: OK"
                    healthy = true
                    break
                } else {
                    echo "Попытка #${i}: ответ ${code}"
                }
            }

            if (!healthy) {
                echo "Ни одна из 5 попыток не вернула 200, откатываемся"
                deployBackup()
                currentBuild.result = 'UNSTABLE'
            } else {
                echo "Хотя бы одна попытка вернула 200 — всё гут"
            }
        }
    }
}


        stage('Test API'){
            steps{
                script{
                    echo "Здесь должно быть тестирование"
                }
            }
        }
    }
}
