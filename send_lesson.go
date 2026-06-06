package main

import (
	"encoding/json"
	"time"
	"net/url"
	"io"
	"net/http"
	"bytes"
	"fmt"
)

type submitData struct {
	lessonData
	CodeLastModAt	string	`json:"codeLastModAt"`
	UserTimezone	string	`json:"userTimezone"`
}

type lessonData struct {
	Input	string		`json:"input"`
	Files	[]fileData	`json:"files"`
}

type fileData struct {
	Name		string	`json:"Name"`
	Content		string	`json:"Content"`
	IsHidden	bool	`json:"IsHidden"`
	IsReadonly	bool	`json:"IsReadonly"`
}

func sendLessons(apiUrl, apiToken string) error {
	payload := submitData{}

	payload.UserTimezone = "Europe/Berlin"

	tz, err := time.LoadLocation(payload.UserTimezone)
	if err != nil {
		return err
	}

	t := time.Now()
	
	payload.CodeLastModAt = t.In(tz).Format(time.RFC3339)

	for lessonID, lessonObject := range Lessons {
		requestURL, err := url.JoinPath(apiUrl, "v1/lessons", lessonID)
		if err != nil {
			return err
		}

		payload.Files = []fileData {
			{
				Name:		"main.py",
				Content:	"# Hacked by Pzykoh",
				IsHidden:	false,
				IsReadonly: false,
			},
		}
		payload.Input = lessonObject.Input

		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(body))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println(string(responseBody))
		time.Sleep(1 * time.Second)
	}
	return nil
}


var Lessons = map[string]lessonData {
	// Learn Python
	"78b4646f-85aa-42c7-ba46-faec2f0902a9": {
		Input: "Welcome to Fantasy Quest!",
	},
	"d3d50474-d60a-4751-b71e-e145ae95b966": {
		Input: "Sam's health is: 100Sam takes 10 damage...Sam's health is: 90",
	},
	"016cddb3-d1d9-4f06-8299-715b7e41ba89": {
		Input: "Greetings, adventurer!",
	},
	"98f52c08-f008-4e2c-a422-73983eecc0c1": {
		Input: "Frontend apps",
	},
	"7c7fcfa3-e8de-4e8e-9f39-d46f7e8420b1": {
		Input: "simple to read and write",
	},
	"45328d3f-5b45-4072-acee-60f9d35386bc": {
		Input: "325",
	},
	"e778983a-b006-4e1e-9797-d4170c31bf80": {
		Input: "Jax: B-Kaw!Hero: ...Jax: Where are you off to this morning? Bkaw...Hero: Where did an owl learn to speak??",
	},
	"366f2438-8f95-4f24-a55c-156a4544b83e": {
		Input: "Welcome to Fantasy Quest!",
	},
	"28607ef6-faa5-42a8-a9bc-ec361b3dd87a": {
		Input: "The rules for valid code in a programming language",
	},
	"470991b7-b331-407c-988d-5ec350690e65": {
		Input: "You'll get an error message from the Python interpreter and the code won't execute",
	},
	"142c8a73-5ede-49a6-9460-563890646023": {
		Input: "Starting up game server...local game server is listening on port 8080",
	},
	"6ba74bb7-3196-4631-90c0-ad01fa3b2ce9": {
		Input: "247.5",
	},
	"b74b2417-979f-4e64-a7bf-acb05cd20bd0": {
		Input: "Ah! Great choices...Is there anything else I can help you with?",
	},
	"7e87cdd3-1dc9-44b0-acc0-a96a40485bc0": {
		Input: "1000",
	},
	"f12268f9-e9a5-4fdc-98d4-444e5cfa0247": {
		Input: "900800700600",
	},
	"c10d6ee8-4cab-47cd-b530-a2b585f349f7": {
		Input: "2000",
	},
	"e0fe3f7c-c0ee-4d02-be2c-b770ddfa477e": {
		Input: "90",
	},
	"9e4e4bf0-6848-4548-af74-887cd9518f63": {
		Input: "scimitar",
	},
	"8a46647a-bd1f-4f2e-8717-4c6df8365128": {
		Input: "heroHealth",
	},
	"24bd9068-1f27-49ab-84c3-a89a39173542": {
		Input: "hero_health",
	},
	"6f85fe3e-c2a4-42e2-b6ca-6d541170d0b1": {
		Input: "player_health is a/an intplayer_has_magic is a/an bool",
	},
	"9c216a9b-6467-4526-88f5-02678df98d77": {
		Input: "Yarl is a dwarf who is 37 years old.",
	},
	"06a78273-7b14-4d5c-b0c1-1883bc065699": {
		Input: "True",
	},
	"70610028-97ba-483e-8ddd-5cbad773dd53": {
		Input: "As the default value that will be replaced later",
	},
	"7d44758a-103c-413a-a594-9d723b99481a": {
		Input: "No",
	},
	"23c2d8ba-9aaa-41d4-87ec-4f8dc8d7ed31": {
		Input: "Dynamic",
	},
	"1ec24fe8-9b50-4e45-9dc2-bfde20704196": {
		Input: "You have 1200 healthYou have 1100 health",
	},
	"9bb5faee-6338-4446-b566-0f84d09508d4": {
		Input: "No limit",
	},
	"88b5f06c-ad26-4e3f-9571-ec6e8f14fec4": {
		Input: "Code that is easy for developers to read and understand",
	},
	"e4fac74c-9d67-41ad-a85c-c579cb3ad76f": {
		Input: "Well met, CharlesThe local mine has been taken over by orcs!We need your help taking it back.",
	},
	"cbd18c5e-236f-497a-a42e-e83810cc0f04": {
		Input: "100",
	},
	"1b5f7530-5d63-41ae-b2bf-d76bf72dce61": {
		Input: "Character ReportLopen is a level 25 Windrunner.They have 15.0 magic resistance.Their account is currently active: True=========================Character Report CompleteData types:name: str, level: int, character_class: strmagic_resistance: floataccount_active: bool",
	},
	"b4f5a4ef-9fb0-40f7-9dc6-baf28fe1be36": {
		Input: "Sword length: 1.0 meters.Sword attack area: 3.14 square metersSpear length: 2.0 meters.Spear attack area: 12.56 square meters",
	},
	"2e076241-9f2e-4efc-9f0f-29b265e2c38b": {
		Input: "The function is defined",
	},
	"ab8febf8-8365-4fb8-99a4-aff41516b6b8": {
		Input: "Because only the value of the variable is passed to the function. It is then assigned to a new variable called 'r'",
	},
	"64266f19-9783-44c1-b63b-01f40ae9a0b8": {
		Input: "Getting damage for 2 4 and 3 ...9 points of damage dealt!=====================================Getting damage for -1 10 and 5 ...14 points of damage dealt!=====================================",
	},
	"cd545883-8994-477c-9412-edd23f1bbdf7": {
		Input: "First name: FrodoLast name: BagginsJob: warriorTitle: Frodo Baggins the warrior=====================================First name: BilboLast name: BagginsJob: thiefTitle: Bilbo Baggins the thief=====================================First name: GandalfLast name: The GreyJob: wizardTitle: Gandalf The Grey the wizard=====================================First name: AragornLast name: Son of ArathornJob: rangerTitle: Aragorn Son of Arathorn the ranger=====================================",
	},
	"c3235f94-174b-49f8-a490-a292f902a63a": {
		Input: "Fantasy Quest is booting up...Game is running!",
	},
	"60741e61-ca0d-4dd5-b1d4-72c6577300ee": {
		Input: "defined, called",
	},
	"c75acfea-7029-4bb3-9609-21f063636c4d": {
		Input: "Create an entrypoint function (usually called `main`) and call it at the end of the file",
	},
	"0f4fa755-1ce7-468b-bf34-c1460e97bf28": {
		Input: "100 degrees fahrenheit is 37.78 degrees celsius88 degrees fahrenheit is 31.11 degrees celsius104 degrees fahrenheit is 40.0 degrees celsius112 degrees fahrenheit is 44.44 degrees celsius",
	},
	"dcafd50a-c53f-4a1a-8677-faf551ce4568": {
		Input: "3 times called and the results are in the stab, slash and fireball damage variables",
	},
	"6b3ce0b8-b323-4685-bf50-48cfd1c9959d": {
		Input: "10 hours is 36000 seconds1 hours is 3600 seconds25 hours is 90000 seconds100 hours is 360000 seconds33 hours is 118800 seconds",
	},
	"6bbc28c7-8887-45ae-85d5-50df81720e35": {
		Input: "It returns None",
	},
	"589229bc-0245-4aca-ac0e-aa07a904b283": {
		Input: "Print: Shows a value in the console.\nReturn: Makes a value available to the caller of the function",
	},
	"3c5fe40f-41e3-4d7e-a035-be67c8d83536": {
		Input: "Frodo Baggins the warrior has a power level of: 6Bilbo Baggins the warrior has a power level of: 11Gandalf The Grey the warrior has a power level of: 9001",
	},
	"5de37882-4ba4-4544-ba0b-ead621b4b451": {
		Input: "Arguments",
	},
	"ddc9202b-8874-480b-8ec7-aec75f0eeaec": {
		Input: "Parameters",
	},
	"3963019a-a8ce-45c0-b988-3e6e4121aa36": {
		Input: "Running tests for health 400 and armor 5========================================Health: 400, Armor: 5Health after punch: 355----------------------------------------Health: 400, Armor: 5Health after slash: 305----------------------------------------Health: 400, Armor: no armor!Health after slash: 300----------------------------------------Health: 400, Armor: no armor!Health after punch: 350----------------------------------------Running tests for health 300 and armor 3========================================Health: 300, Armor: 3Health after punch: 253----------------------------------------Health: 300, Armor: 3Health after slash: 203----------------------------------------Health: 300, Armor: no armor!Health after slash: 200----------------------------------------Health: 300, Armor: no armor!Health after punch: 250----------------------------------------Running tests for health 200 and armor 1========================================Health: 200, Armor: 1Health after punch: 151----------------------------------------Health: 200, Armor: 1Health after slash: 101----------------------------------------Health: 200, Armor: no armor!Health after slash: 100----------------------------------------Health: 200, Armor: no armor!Health after punch: 150----------------------------------------",
	},
	"cb3a6f03-90e6-46d7-8647-6c271c05a108": {
		Input: "Weapon's base damage: 100.0Cursing...With lesser curse the damage is: 50.0 damage.With greater curse the damage is: 25.0 damage.=====================================Weapon's base damage: 500.0Cursing...With lesser curse the damage is: 250.0 damage.With greater curse the damage is: 125.0 damage.=====================================Weapon's base damage: 1000.0Cursing...With lesser curse the damage is: 500.0 damage.With greater curse the damage is: 250.0 damage.=====================================",
	},
	"c6956acb-a130-4aef-9907-c4a5eb601f36": {
		Input: "A place to practice what you've already learned with personalized challenges before forging ahead in the course",
	},
	"7d07216c-d599-40c2-976e-b87027dd5f12": {
		Input: "The target has 100 health.sword base damage: 50... Enchanting and attacking.The target has been attacked with the enchanted sword.The target has 40 health remaining.=====================================The target has 500 health.axe base damage: 100... Enchanting and attacking.The target has been attacked with the enchanted axe.The target has 390 health remaining.=====================================The target has 1000 health.bow base damage: 250... Enchanting and attacking.The target has been attacked with the enchanted bow.The target has 740 health remaining.=====================================",
	},
	"4777c0b2-30fa-48fe-82bf-c9b84e74d92f": {
		Input: "Archmage",
	},
	// First Personal Project
	"b34b0f83-0af0-4bad-9e8d-65ebcd8d7cbc": {
		Input: "Yes",
	},
	"e8996535-f3cd-4a96-b525-c31afc7ae7b7": {
		Input: "Yes",
	},
	"3b4d8418-1c88-4d41-a135-bf3acbc04ce3": {
		Input: "Yes",
	},
	"d4dc954b-06cf-4f7b-bc36-fa9936c245ec": {
		Input: "https://github.com/luho91/bootdev-fellowship",
	},
	// Second Personal Project
	"73e3f3eb-c845-43c2-acfe-717e6fa8590a": {
		Input: "Yes",
	},
	"71db2789-70c3-484d-a5e0-8e5c0f6b2b30": {
		Input: "Yes",
	},
	"ab815d5f-f878-4e93-85e1-eaa9ae41d626": {
		Input: "Yes",
	},
	"28316b24-fa39-4172-b2f6-4ea46ff1cebc": {
		Input: "https://github.com/luho91/bootdev-fellowship",
	},
	"f2e60d86-6956-41b0-be1f-a804b6fbe2f6": {
		Input: "Yes",
	},
	"5e864ce9-c485-4b9e-9bf5-f0692f578050": {
		Input: "Yes",
	},
	"dff59594-9152-480d-85b8-45da1fc7fc68": {
		Input: "https://github.com/luho91/bootdev-fellowship",
	},
}
